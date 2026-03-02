package processors

import (
	"bufio"
	"fmt"
	"io"
)

type StreamingMode string

const (
	StreamingModeNone     StreamingMode = "none"
	StreamingModeBuffered StreamingMode = "buffered"
	StreamingModeChunked  StreamingMode = "chunked"
	StreamingModeLine     StreamingMode = "line"
	StreamingModeNative   StreamingMode = "native"
)

type StreamingSpec struct {
	Mode      StreamingMode
	ChunkSize int
	Prefer    bool
}

type StreamConfigurable interface {
	Processor
	StreamingSpec() StreamingSpec
}

type NativeStreamProcessor interface {
	Processor
	TransformStream(reader io.Reader, writer io.Writer, opts ...Flag) error
}

var DefaultStreamingSpec = StreamingSpec{
	Mode:      StreamingModeNone,
	ChunkSize: 64 * 1024,
	Prefer:    false,
}

func normalizeStreamingSpec(spec StreamingSpec) StreamingSpec {
	if spec.Mode == "" {
		spec.Mode = StreamingModeNone
	}
	if spec.ChunkSize <= 0 {
		spec.ChunkSize = DefaultStreamingSpec.ChunkSize
	}
	return spec
}

func GetStreamingSpec(processor Processor) StreamingSpec {
	if sp, ok := processor.(StreamConfigurable); ok {
		spec := normalizeStreamingSpec(sp.StreamingSpec())
		if spec.Mode == StreamingModeNative {
			if _, ok := processor.(NativeStreamProcessor); ok {
				return spec
			}
			spec.Mode = StreamingModeNone
		}
		return spec
	}
	return DefaultStreamingSpec
}

func SupportsStreaming(processor Processor) bool {
	spec := GetStreamingSpec(processor)
	return spec.Mode != StreamingModeNone
}

func ShouldStream(processor Processor, inputSize, threshold int64) bool {
	spec := GetStreamingSpec(processor)
	if spec.Mode == StreamingModeNone {
		return false
	}
	if spec.Prefer {
		return true
	}
	return inputSize > threshold
}

func TransformStream(processor Processor, reader io.Reader, writer io.Writer, opts ...Flag) error {
	spec := GetStreamingSpec(processor)
	spec = normalizeStreamingSpec(spec)

	switch spec.Mode {
	case StreamingModeNative:
		nsp, ok := processor.(NativeStreamProcessor)
		if !ok {
			return fmt.Errorf("processor %q declares native streaming without TransformStream", processor.Name())
		}
		return nsp.TransformStream(reader, writer, opts...)
	case StreamingModeLine:
		return transformStreamLineByLine(processor, reader, writer, opts...)
	case StreamingModeBuffered:
		return transformStreamBuffered(processor, reader, writer, opts...)
	case StreamingModeChunked:
		return transformStreamChunked(processor, reader, writer, spec.ChunkSize, opts...)
	default:
		return fmt.Errorf("processor %q does not support streaming", processor.Name())
	}
}

func transformStreamLineByLine(p Processor, r io.Reader, w io.Writer, opts ...Flag) error {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		result, err := p.Transform(scanner.Bytes(), opts...)
		if err != nil {
			return err
		}
		if _, err := fmt.Fprintln(w, result); err != nil {
			return err
		}
	}
	return scanner.Err()
}

func transformStreamBuffered(p Processor, r io.Reader, w io.Writer, opts ...Flag) error {
	data, err := io.ReadAll(r)
	if err != nil {
		return err
	}
	result, err := p.Transform(data, opts...)
	if err != nil {
		return err
	}
	_, err = w.Write([]byte(result))
	return err
}

func transformStreamChunked(p Processor, r io.Reader, w io.Writer, chunkSize int, opts ...Flag) error {
	buf := make([]byte, chunkSize)
	for {
		n, err := r.Read(buf)
		if n > 0 {
			result, tErr := p.Transform(buf[:n], opts...)
			if tErr != nil {
				return tErr
			}
			if _, wErr := w.Write([]byte(result)); wErr != nil {
				return wErr
			}
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
	}
	return nil
}

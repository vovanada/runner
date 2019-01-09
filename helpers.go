package runner

import "io"

func captureOut(r io.Reader, out *[]byte) error {
	buf := make([]byte, 1024, 1024)
	for {
		n, err := r.Read(buf[:])
		if n > 0 {
			d := buf[:n]
			*out = append(*out, d...)
		}
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			return err
		}
	}
}

package utils

import (
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

func DiffFileHeadTail(filename1 string, filename2 string) error {
	const diffsize = 2048

	var f1 *os.File
	var f2 *os.File
	var f1info os.FileInfo
	var f2info os.FileInfo
	var err error
	hex2string := hex.EncodeToString

	{
		f1, err = os.Open(filename1)
		if err != nil {
			return err
		}
		f1info, err = f1.Stat()
		if err != nil {
			return err
		}
	}

	{
		f2, err = os.Open(filename2)
		if err != nil {
			return err
		}
		f2info, err = f2.Stat()
		if err != nil {
			return err
		}
	}

	if f1info.Size() != f2info.Size() {
		return fmt.Errorf("size not equal: %s: %d, %s: %d", filename1, f1info.Size(), filename2, f2info.Size())
	}

	getheadtail := func(f *os.File, fs os.FileInfo) (head []byte, tail []byte, err error) {
		// head := make([]byte, 16)
		// tail := make([]byte, 16)
		if fs.Size() < diffsize {
			buf, err := io.ReadAll(f)
			if err != nil {
				return nil, nil, err
			}
			return buf, buf, nil
		}

		head = make([]byte, diffsize)
		tail = make([]byte, diffsize)

		_, err = io.ReadFull(f, head)
		if err != nil {
			return nil, nil, err
		}

		_, err = f.Seek(-diffsize, 2)
		if err != nil {
			return nil, nil, err
		}
		_, err = io.ReadFull(f, tail)
		if err != nil {
			return nil, nil, err
		}
		return head, tail, nil
	}

	head1, tail1, err := getheadtail(f1, f1info)
	if err != nil {
		return err
	}

	head2, tail2, err := getheadtail(f2, f1info)
	if err != nil {
		return err
	}

	if hex2string(head1) != hex2string(head2) {
		return fmt.Errorf("head not equal\n%s: %s\n%s: %s", filename1, hex2string(head1), filename2, hex2string(head2))
	}

	if hex2string(tail1) != hex2string(tail2) {
		return fmt.Errorf("head not equal\n%s: %s\n%s: %s", filename1, hex2string(head1), filename2, hex2string(head2))
	}

	return nil
}

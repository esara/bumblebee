// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/bumblebee/api/bumblebee.io/probes/v1alpha1/probe.proto

package v1alpha1

import (
	"encoding/binary"
	"errors"
	"fmt"
	"hash"
	"hash/fnv"

	"github.com/mitchellh/hashstructure"
	safe_hasher "github.com/solo-io/protoc-gen-ext/pkg/hasher"
)

// ensure the imports are used
var (
	_ = errors.New("")
	_ = fmt.Print
	_ = binary.LittleEndian
	_ = new(hash.Hash64)
	_ = fnv.New64
	_ = hashstructure.Hash
	_ = new(safe_hasher.SafeHasher)
)

// Hash function
func (m *ProbeSpec) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("probes.bumblebee.io.github.com/solo-io/bumblebee/pkg/api/probes.bumblebee.io/v1alpha1.ProbeSpec")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetImage())); err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetImagePullPolicy())
	if err != nil {
		return 0, err
	}

	{
		var result uint64
		innerHash := fnv.New64()
		for k, v := range m.GetNodeSelector() {
			innerHash.Reset()

			if _, err = innerHash.Write([]byte(v)); err != nil {
				return 0, err
			}

			if _, err = innerHash.Write([]byte(k)); err != nil {
				return 0, err
			}

			result = result ^ innerHash.Sum64()
		}
		err = binary.Write(hasher, binary.LittleEndian, result)
		if err != nil {
			return 0, err
		}

	}

	{
		var result uint64
		innerHash := fnv.New64()
		for k, v := range m.GetAdditionalLabels() {
			innerHash.Reset()

			if _, err = innerHash.Write([]byte(v)); err != nil {
				return 0, err
			}

			if _, err = innerHash.Write([]byte(k)); err != nil {
				return 0, err
			}

			result = result ^ innerHash.Sum64()
		}
		err = binary.Write(hasher, binary.LittleEndian, result)
		if err != nil {
			return 0, err
		}

	}

	return hasher.Sum64(), nil
}

package estr_test

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/wusphinx/estr"
)

type EStrTestSuite struct {
	suite.Suite
	testcases []string
}

func (s *EStrTestSuite) SetupTest() {
	s.testcases = []string{
		"你好！世界",
		"https://www.google.com.hk/",
		"hello！world",
		"hello！world👀",
	}
}

func (s *EStrTestSuite) TestDefault() {
	err := estr.Init("", 0)
	s.Nil(err)

	for _, src := range s.testcases {
		dst, err := estr.Encode(src)
		s.Nil(err)
		s.T().Logf("dst is: %s", dst)

		res, err := estr.Decode(dst)
		s.Nil(err)

		s.Equal(src, res)
	}
}

func (s *EStrTestSuite) TestSet() {
	settings := map[string]uint8{
		"a": 10,
		"人": 10,
		"👀": 100,
	}

	for salt, minLength := range settings {
		err := estr.Init(salt, minLength)
		s.Nil(err)

		for _, src := range s.testcases {
			dst, err := estr.Encode(src)
			s.Nil(err)
			s.T().Logf("dst is: %s", dst)

			res, err := estr.Decode(dst)
			s.Nil(err)

			s.Equal(src, res)
		}
	}
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(EStrTestSuite))
}

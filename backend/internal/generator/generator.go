package generator

import "k6clone/internal/model"

type Generator interface {
	Generate(url string) (*model.Script, error)
}
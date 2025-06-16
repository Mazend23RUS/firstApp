package implementationUseCase

type BuildUseCase struct {
	userUseCase *AuthUseCase
}

func NewBuilder() *BuildUseCase {
	return &BuildUseCase{
		userUseCase: &AuthUseCase{},
	}
}

// func (b *BuildUseCase) Build() *AuthUseCase {

// }

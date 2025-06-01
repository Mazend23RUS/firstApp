package user

import userDTO "github.com/alexey/boundary/dto"

type CreateRequest struct {
	Name        string `json:"name" validate:"required,min=2,max=50"`
	Email       string `json:"email" validate:"required,email"`
	Password    string `json:"password" validate:"required,min=6"`
	PhoneNumber int    `json:"phone" validate:"required,min=12,max=12"`
}

type userResponse struct {
}




// type MapperOfRequestToDTO struct {
//    CreateRequest

// }

func MapperOfRequestToDTO(c *CreateRequest) userDTO.UserDTO {
    user := userDTO.UserDTO{
		PhoneNumber : c.PhoneNumber,
		Email: c.Email,
        Name: c.Name,
		Password: c.Password,
	}

   return user
}
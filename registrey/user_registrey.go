package registry

import (
	"github.com/VooDooStack/FitStackAPI/interface/controller"
	ip "github.com/VooDooStack/FitStackAPI/interface/presenter"
	ir "github.com/VooDooStack/FitStackAPI/interface/repository"
	"github.com/VooDooStack/FitStackAPI/usecase/interactor"
	up "github.com/VooDooStack/FitStackAPI/usecase/presenter"
	ur "github.com/VooDooStack/FitStackAPI/usecase/repository"
)

func (r *registry) NewUserController() controller.UserController {
	return controller.NewUserController(r.NewUserInteractor())
}

func (r *registry) NewUserInteractor() interactor.UserInteractor {
	return interactor.NewUserInteractor(r.NewUserRepository(), r.NewUserPresenter())
}

func (r *registry) NewUserRepository() ur.UserRepository {
	return ir.NewUserRepository(r.db)
}

func (r *registry) NewUserPresenter() up.UserPresenter {
	return ip.NewUserPresenter()
}

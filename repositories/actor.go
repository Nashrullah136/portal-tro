package repositories

import (
	"context"
	"gorm.io/gorm"
	"nashrul-be/crm/entities"
)

type ActorRepositoryInterface interface {
	GetAll(ctx context.Context, username, role string, limit, offset uint) (actor []entities.User, err error)
	GetByUsername(ctx context.Context, username string) (actor entities.User, err error)
	GetByUsernameBatch(ctx context.Context, username []string) (actors []entities.User, err error)
	IsUsernameExist(actor entities.User) (exist bool, err error)
	Create(ctx context.Context, actor entities.User) (result entities.User, err error)
	Update(ctx context.Context, actor entities.User) (result entities.User, err error)
	Save(ctx context.Context, actor entities.User) (result entities.User, err error)
	Delete(ctx context.Context, username string) (err error)
}

func NewActorRepository(db *gorm.DB) ActorRepositoryInterface {
	return actorRepository{db: db}
}

type actorRepository struct {
	db *gorm.DB
}

func (r actorRepository) GetAll(ctx context.Context, username, role string, limit, offset uint) (actor []entities.User, err error) {
	query := r.db.WithContext(ctx).Model(&entities.User{}).Preload("Role").Where("username LIKE ?", username).
		Limit(int(limit)).Offset(int(offset))
	if role != "" {
		switch role {
		case "admin":
			query = query.Where("role_id = ?", 1)
		case "user":
			query = query.Where("role_id = ?", 2)
		}
	}
	err = query.Find(&actor).Error
	return
}

func (r actorRepository) GetByUsername(ctx context.Context, username string) (actor entities.User, err error) {
	actor.Username = username
	err = r.db.WithContext(ctx).First(&actor).Error
	return
}

func (r actorRepository) GetByUsernameBatch(ctx context.Context, username []string) (actors []entities.User, err error) {
	err = r.db.WithContext(ctx).Find(&actors, "username IN ?", username).Error
	return
}

func (r actorRepository) IsUsernameExist(actor entities.User) (exist bool, err error) {
	var count int64
	err = r.db.Model(&entities.User{}).Where("username = ?", actor.Username).Count(&count).Error
	if err != nil {
		return
	}
	exist = count > 0
	return
}

func (r actorRepository) Create(ctx context.Context, actor entities.User) (result entities.User, err error) {
	result = actor
	err = r.db.WithContext(ctx).Create(&result).Error
	return
}

func (r actorRepository) Update(ctx context.Context, actor entities.User) (result entities.User, err error) {
	result = actor
	err = r.db.WithContext(ctx).Updates(&result).Error
	return
}

func (r actorRepository) Save(ctx context.Context, actor entities.User) (result entities.User, err error) {
	result = actor
	err = r.db.WithContext(ctx).Save(&result).Error
	return
}

func (r actorRepository) Delete(ctx context.Context, username string) (err error) {
	user := entities.User{
		Username: username,
	}
	err = r.db.WithContext(ctx).Delete(&user).Error
	return
}

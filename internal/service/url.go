package service

import (
	"errors"
	"fmt"

	"github.com/Lacky1234union/UrlShorter/internal/lib/random"
	"github.com/Lacky1234union/UrlShorter/internal/storage"
)

// URLService defines the interface for URL shortening operations
type URLService interface {
	SaveURL(url string, alias string) (string, error)
	GetURL(alias string) (string, error)
	DeleteURL(alias string) error
}

// urlService implements URLService interface
type urlService struct {
	storage storage.URLStorage
}

// NewURLService creates a new instance of URLService
func NewURLService(storage storage.URLStorage) URLService {
	return &urlService{
		storage: storage,
	}
}

// SaveURL saves a URL with the given alias or generates a random one
func (s *urlService) SaveURL(url string, alias string) (string, error) {
	if url == "" {
		return "", errors.New("url cannot be empty")
	}

	if alias == "" {
		alias = random.NewRandomString(6)
	}

	id, err := s.storage.SaveURL(url, alias)
	if err != nil {
		return "", fmt.Errorf("failed to save URL: %w", err)
	}

	return alias, nil
}

// GetURL retrieves the original URL for the given alias
func (s *urlService) GetURL(alias string) (string, error) {
	if alias == "" {
		return "", errors.New("alias cannot be empty")
	}

	url, err := s.storage.GetURL(alias)
	if err != nil {
		return "", fmt.Errorf("failed to get URL: %w", err)
	}

	return url, nil
}

// DeleteURL deletes the URL with the given alias
func (s *urlService) DeleteURL(alias string) error {
	if alias == "" {
		return errors.New("alias cannot be empty")
	}

	err := s.storage.DeleteURL(alias)
	if err != nil {
		return fmt.Errorf("failed to delete URL: %w", err)
	}

	return nil
}

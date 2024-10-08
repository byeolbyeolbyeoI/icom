package repository

import "github.com/byeolbyeolbyeoI/widyanaya-api/internal/publication/model"

type PublicationRepositoryInstance interface {
	IsPublisherExists(int) error

	IsPublicationsExist() error
	IsPublicationExists(int) error
	IsPublicationsExistByCategoryId(int) error
	IsPublicationCategoryExists(int) error

	GetPublications() ([]model.Publication, error)
	GetPublicationsByCategoryId(int) ([]model.Publication, error)
	GetPublicationById(int) (model.Publication, error)
	AddPublication(model.Publication) error
	UpdatePublication(model.UpdatedPublication) error
	DeletePublicationById(int) error

	IsPaperExists(int) error
	IsPapersExist() error
	IsOwnerExists(int) error
	GetPapers() ([]model.Paper, error)
	GetPaperById(int) (model.Paper, error)
	AddPaper(model.Paper) error
	UpdatePaper(model.UpdatedPaper) error
	DeletePaperById(int) error

	IsPaperFragmentsExist() error
	GetPaperFragments() ([]model.PaperFragment, error)
	GetPaperFragmentById(int) (model.PaperFragment, error)
	IsPaperFragmentExists(int) error
	IsPaperFragmentCategoryExists(int) error
	AddPaperFragment(model.PaperFragment) error
	UpdatePaperFragment(model.UpdatedPaperFragment) error
	DeletePaperFragmentById(int) error

	IsCompetitionExists(int) error
	IsCompetitionCategoryExists(int) error
	IsCompetitionsExist() error
	IsCompetitionsExistByCategoryId(int) error
	GetCompetitions() ([]model.Competition, error)
	GetCompetitionById(int) (model.Competition, error)
	GetCompetitionsByCategoryId(int) ([]model.Competition, error)
	AddCompetition(model.Competition) error
	UpdateCompetition(model.UpdatedCompetition) error
	DeleteCompetitionById(int) error

	IsPublicationRequestExists(int) error
	IsPublicationRequestsExist() error
	IsReferenceFormatExists(int) error
	IsRequesterExists(int) error
	GetPublicationRequests() ([]model.PublicationRequest, error)
	GetPublicationRequestById(int) (model.PublicationRequest, error)
	AddPublicationRequest(model.PublicationRequest) error
	UpdatePublicationRequest(model.UpdatedPublicationRequest) error
	DeletePublicationRequestById(int) error

	IsMetadataExists(int) error
	IsMetadatasExist() error

	GetMetadatas() ([]model.Metadata, error)
	GetMetadataById(int) (model.Metadata, error)
	AddMetadata(model.Metadata) error
	UpdateMetadata(model.UpdatedMetadata) error
	DeleteMetadataById(int) error
}

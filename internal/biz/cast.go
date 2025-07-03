package biz

type MediaAccountRepo interface {
}

type WorkRepo interface {
}

type WorkImageRepo interface {
}

type WorkVideoRepo interface {
}

type WorkLogRepo interface {
}

type CostLogRepo interface {
}

type BundleOrderRecordsRepo interface {
}

type CastUseCase struct {
	mediaAccRepo           MediaAccountRepo
	workRepo               WorkRepo
	workImageRepo          WorkImageRepo
	workVideoRepo          WorkVideoRepo
	workLogRepo            WorkLogRepo
	costLogRepo            CostLogRepo
	bundleOrderRecordsRepo BundleOrderRecordsRepo
}

func NewCastUseCase(mediaAccRepo MediaAccountRepo, workRepo WorkRepo, workImageRepo WorkImageRepo,
	workVideoRepo WorkVideoRepo, workLogRepo WorkLogRepo, costLogRepo CostLogRepo, bundleOrderRecordsRepo BundleOrderRecordsRepo) *CastUseCase {
	return &CastUseCase{mediaAccRepo: mediaAccRepo, workRepo: workRepo, workImageRepo: workImageRepo,
		workVideoRepo: workVideoRepo, workLogRepo: workLogRepo, costLogRepo: costLogRepo, bundleOrderRecordsRepo: bundleOrderRecordsRepo}
}

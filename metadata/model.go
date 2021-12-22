package metadata

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/metapass-metadata/config"
	"github.com/metapass-metadata/db"
)

const METAPASS_IMAGE_URL string = "https://storage.googleapis.com/metapass-images/"
const EXTERNAL_URL string = "https://enterdao.xyz/sharded-mind/"
const GENES_COUNT = 7

const SKIN_GENE_COUNT int = 11
const MOUTH_GENE_COUNT int = 15
const NECKLACES_GENE_COUNT int = 16
const EYES_GENES_COUNT int = 17
const TRACKS_GENES_COUNT int = 10
const BACKGROUND_GENE_COUNT int = 18
const VORTEX_GENE_COUNT = 19

var SKIN_DISITRIBUTION [SKIN_GENE_COUNT]int = [SKIN_GENE_COUNT]int{
	400,
	400,
	480,
	600,
	1007,
	1007,
	1009,
	1009,
	1201,
	1201,
	1686,
}
var MOUTH_DISITRIBUTION [MOUTH_GENE_COUNT]int = [MOUTH_GENE_COUNT]int{
	100,
	186,
	480,
	550,
	550,
	600,
	600,
	600,
	600,
	700,
	801,
	801,
	841,
	1001,
	1590,
}
var NECKLACES_DISTRIBUTION [NECKLACES_GENE_COUNT]int = [NECKLACES_GENE_COUNT]int{
	100,
	198,
	300,
	494,
	494,
	500,
	500,
	600,
	700,
	700,
	800,
	800,
	800,
	901,
	901,
	1201,
}
var EYES_DISTRIBUTION [EYES_GENES_COUNT]int = [EYES_GENES_COUNT]int{
	100,
	200,
	300,
	400,
	446,
	480,
	500,
	500,
	600,
	700,
	700,
	700,
	700,
	700,
	741,
	841,
	1391,
}
var TRACKS_DISTRIBUTION [TRACKS_GENES_COUNT]int = [TRACKS_GENES_COUNT]int{
	86,
	300,
	500,
	720,
	901,
	1101,
	1321,
	1501,
	1682,
	1888,
}
var BACKGROUND_DISTRIBUTION [BACKGROUND_GENE_COUNT]int = [BACKGROUND_GENE_COUNT]int{
	96,
	98,
	98,
	298,
	298,
	298,
	500,
	500,
	500,
	500,
	700,
	700,
	700,
	700,
	1001,
	1001,
	1001,
	1011,
}
var VORTEX_DISTRIBUTION [VORTEX_GENE_COUNT]int = [VORTEX_GENE_COUNT]int{
	194,
	194,
	398,
	400,
	400,
	400,
	500,
	500,
	500,
	500,
	500,
	600,
	600,
	600,
	600,
	700,
	800,
	800,
	800,
}

const backgroundsFolder = "backgrounds-image"
const skinsFolder = "skins"
const necklacesFolder = "necklaces"
const mouthFolder = "mouth"
const eyesFolder = "eyes"
const vortexFolder = "vortex"
const tracksFolder = "tracks"

type Genome string
type Gene int
type StringAttribute struct {
	TraitType string `json:"trait_type"`
	Value     string `json:"value"`
}

type IntegerAttribute struct {
	TraitType string `json:"trait_type"`
	Value     int    `json:"value"`
}

type FloatAttribute struct {
	TraitType   string  `json:"trait_type"`
	Value       float64 `json:"value"`
	DisplayType string  `json:"display_type"`
}

func (g Gene) toPath() string {
	return strconv.Itoa(int(g))
}

func getGeneInt(g string, start, end int, traitDistributions []int) int {
	genomeLen := len(g)
	geneStr := g[genomeLen+start : genomeLen+end]
	gene, _ := strconv.Atoi(geneStr)
	counter, traitIdx := 0, 0

	for i, distribution := range traitDistributions {
		counter += distribution
		if gene <= counter {
			traitIdx = i
			break
		}
	}

	return traitIdx
}

// BACKGROUND_DISTRIBUTION
func getBackgroundGene(g string) int {
	return getGeneInt(g, -4, 0, BACKGROUND_DISTRIBUTION[:])
}
// SKIN_DISITRIBUTION
func getSkinGene(g string) int {
	return getGeneInt(g, -8, -4, SKIN_DISITRIBUTION[:])
}
// NECKLACES_DISTRIBUTION
func getNecklaceGene(g string) int {
	return getGeneInt(g, -12, -8, NECKLACES_DISTRIBUTION[:])
}
// MOUTH_DISITRIBUTION
func getMouthGene(g string) int {
	return getGeneInt(g, -16, -12, MOUTH_DISITRIBUTION[:])
}
// EYES_DISTRIBUTION
func getEyesGene(g string) int {
	return getGeneInt(g, -20, -16, EYES_DISTRIBUTION[:])
}
// VORTEX_DISTRIBUTION
func getVortexGene(g string) int {
	return getGeneInt(g, -24, -20, VORTEX_DISTRIBUTION[:])
}
// TRACKS_DISTRIBUTION
func getTrackGene(g string) int {
	return getGeneInt(g, -28, -24, TRACKS_DISTRIBUTION[:])
}

func getMouthGeneAttribute(g string, configService *config.ConfigService) StringAttribute {
	gene := getMouthGene(g)
	return StringAttribute{
		TraitType: "Mouth",
		Value:     configService.Mouth[gene],
	}
}

func getTrackGeneAttribute(g string, configService *config.ConfigService) StringAttribute {
	gene := getTrackGene(g)
	return StringAttribute{
		TraitType: "Track",
		Value:     configService.Tracks[gene],
	}
}

func getVortexGeneAttribute(g string, configService *config.ConfigService) StringAttribute {
	gene := getVortexGene(g)
	return StringAttribute{
		TraitType: "Vortex",
		Value:     configService.Vortex[gene],
	}
}

func getNecklaceGeneAttribute(g string, configService *config.ConfigService) StringAttribute {
	gene := getNecklaceGene(g)
	return StringAttribute{
		TraitType: "Necklace",
		Value:     configService.Necklaces[gene],
	}
}

func getEyesGeneAttribute(g string, configService *config.ConfigService) StringAttribute {
	gene := getEyesGene(g)
	return StringAttribute{
		TraitType: "Eyes",
		Value:     configService.Eyes[gene],
	}
}
func getSkinGeneAttribute(g string, configService *config.ConfigService) StringAttribute {
	gene := getSkinGene(g)
	return StringAttribute{
		TraitType: "Skin",
		Value:     configService.Skins[gene],
	}
}
func getBackgroundGeneAttribute(g string, configService *config.ConfigService) StringAttribute {
	gene := getBackgroundGene(g)
	return StringAttribute{
		TraitType: "Background",
		Value:     configService.Backgrounds[gene],
	}
}

func getEyesGenePath(g string) string {
	gene := getEyesGene(g)
	return Gene(gene).toPath()
}

func getSkinGenePath(g string) string {
	gene := getSkinGene(g)
	return Gene(gene).toPath()
}

func getMouthGenePath(g string) string {
	gene := getMouthGene(g)
	return Gene(gene).toPath()
}

func getNecklaceGenePath(g string) string {
	gene := getNecklaceGene(g)
	return Gene(gene).toPath()
}

func getTrackGenePath(g string) string {
	gene := getTrackGene(g)
	return Gene(gene).toPath()
}

func getVortexGenePath(g string) string {
	gene := getVortexGene(g)
	return Gene(gene).toPath()
}
func getBackgroundGenePath(g string) string {
	gene := getBackgroundGene(g)
	return Gene(gene).toPath()
}

func (g *Genome) name(configService *config.ConfigService, tokenId string) string {
	return fmt.Sprintf("EnterDAO Sharded Minds #%v", tokenId)
}

func (g *Genome) description(configService *config.ConfigService, tokenId string) string {
	return "EnterDAO Sharded Minds is a collection of 5,000 audiovisual art-pieces created by Angela Pencheva and Raredub as contributors to EnterDAO"
}

func (g *Genome) genes() []string {
	gStr := string(*g)

	res := make([]string, 0, GENES_COUNT)

	res = append(res, getBackgroundGenePath(gStr))
	res = append(res, getSkinGenePath(gStr))
	res = append(res, getNecklaceGenePath(gStr))
	res = append(res, getMouthGenePath(gStr))
	res = append(res, getEyesGenePath(gStr))
	res = append(res, getVortexGenePath(gStr))
	res = append(res, getTrackGenePath(gStr))

	return res
}

func (g *Genome) attributes(configService *config.ConfigService) []interface{} {
	gStr := string(*g)

	res := []interface{}{}
	res = append(res, getTrackGeneAttribute(gStr, configService))
	res = append(res, getVortexGeneAttribute(gStr, configService))
	res = append(res, getEyesGeneAttribute(gStr, configService))
	res = append(res, getMouthGeneAttribute(gStr, configService))
	res = append(res, getNecklaceGeneAttribute(gStr, configService))
	res = append(res, getSkinGeneAttribute(gStr, configService))
	res = append(res, getBackgroundGeneAttribute(gStr, configService))
	return res
}

// TODO: Check overlay order (https://enterdao.slack.com/archives/C02CD2N5TBK/p1638837850047000)
func (g *Genome) Metadata(tokenId string, configService *config.ConfigService) Metadata {
	var m Metadata
	genes := g.genes()
	m.Attributes = g.attributes(configService)
	m.Name = g.name(configService, tokenId)
	m.Description = g.description(configService, tokenId)
	m.ExternalUrl = fmt.Sprintf("%s%s", EXTERNAL_URL, tokenId)

	b := strings.Builder{}

	b.WriteString(METAPASS_IMAGE_URL) // Start with base url

	for _, gene := range genes {
		b.WriteString(gene)
	}

	geneUrl := b.String()

	m.Image = geneUrl + ".jpg"
	m.AnimationUrl = geneUrl + ".mp4"

	forVideoImageUrl := fmt.Sprintf("%s-for-video.png", geneUrl)

	dbName := os.Getenv("QUEUE_DB")
	queueCollection := os.Getenv("QUEUE_COLLECTION")
	
	videoImageExists  := resourceExists(forVideoImageUrl)
	imageExists := resourceExists(fmt.Sprintf("%s.jpg", geneUrl))

	videoExists := db.CheckVideoIsQueuedForGeneration(dbName, queueCollection, tokenId)

	if !imageExists {
		GenerateAndSaveImage(genes)
	}
	if !videoImageExists {
		GenerateAndSaveImageForVideo(genes)		
	}

	if !videoExists {
		videoUrl := fmt.Sprintf("%s/backgrounds-video/%s.mp4", BUCKET_BASE_PATH, genes[0])
		trackUrl := fmt.Sprintf("%s/tracks/%s.wav", BUCKET_BASE_PATH, genes[len(genes) - 1])
		forVideoUrl := fmt.Sprintf("%s-for-video.png", strings.Join(genes[:], ""))
		videoName := fmt.Sprintf("%s.mp4", strings.Join(genes[:], ""))
		db.AddVideoForGeneration(db.VideoMetadata{
			TokenId:        tokenId,
			TrackUrl:       trackUrl,
			BackgroundUrl:  videoUrl,
			ImageUrl:       forVideoUrl,
			FinalVideoName: videoName,
			IsGenerated:    false,
			IsGenerating:   false,
			HasErrored:     false,
		})
	}

	return m
}

type Metadata struct {
	Description  string      `json:"description"`
	Name         string      `json:"name"`
	Image        string      `json:"image"`
	AnimationUrl string      `json:"animation_url"`
	Attributes   interface{} `json:"attributes"`
	ExternalUrl  string      `json:"external_url"`
}

package metadata

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/lobster-metadata/config"
)

const METAPASS_IMAGE_URL string = "https://storage.googleapis.com/metapass-images/"
const EXTERNAL_URL string = "https://enterdao.xyz/metapass/"
const GENES_COUNT = 9
const BACKGROUND_GENE_COUNT int = 8
const MOUTH_GENES_COUNT int = 17
const SKIN_GENES_COUNT int = 18
const HAND_GENES_COUNT int = 25
const CLOTHES_GENES_COUNT int = 20
const EYES_GENES_COUNT int = 23
const HEAD_GENES_COUNT int = 22

var SKIN_DISITRIBUTION [SKIN_GENES_COUNT]int = [SKIN_GENES_COUNT]int{1750, 1750, 1750, 1500, 750, 500, 400, 350, 300, 250, 175, 150, 125, 100, 65, 50, 25, 10}
var CLOTHES_DISTRIBUTION [CLOTHES_GENES_COUNT]int = [CLOTHES_GENES_COUNT]int{1750, 1750, 1500, 1000, 850, 500, 450, 400, 350, 300, 250, 200, 175, 150, 125, 100, 65, 50, 25, 10}
var HAND_DISTRIBUTION [HAND_GENES_COUNT]int = [HAND_GENES_COUNT]int{1150, 1000, 1000, 850, 750, 650, 600, 550, 500, 450, 400, 350, 300, 250, 225, 200, 175, 150, 125, 100, 75, 65, 50, 25, 10}
var MOUTH_DISTRIBUTION [MOUTH_GENES_COUNT]int = [MOUTH_GENES_COUNT]int{1350, 1250, 1150, 1050, 950, 850, 750, 650, 550, 450, 350, 250, 150, 100, 75, 50, 25}
var EYES_DISTRIBUTION [EYES_GENES_COUNT]int = [EYES_GENES_COUNT]int{1400, 1300, 1200, 1000, 850, 700, 600, 500, 450, 400, 350, 250, 200, 175, 150, 125, 100, 75, 55, 45, 35, 25, 15}
var HEAD_DISTRIBUTION [HEAD_GENES_COUNT]int = [HEAD_GENES_COUNT]int{750, 600, 600, 600, 600, 600, 600, 600, 600, 600, 600, 600, 600, 550, 400, 325, 250, 200, 150, 100, 50, 25}
var BACKGROUND_DISTRIBUTION [BACKGROUND_GENE_COUNT]int = [BACKGROUND_GENE_COUNT]int{1250, 1250, 1250, 1250, 1250, 1250, 1250, 1250}

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
	if g < 10 {
		return fmt.Sprintf("0%s", strconv.Itoa(int(g)))
	}

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
func getBackgroundGene(g string) int {
	return getGeneInt(g, -4, 0, BACKGROUND_DISTRIBUTION[:])
}
func getSkinGene(g string) int {
	return getGeneInt(g, -8, -4, SKIN_DISITRIBUTION[:])
}
func getClothesGene(g string) int {
	return getGeneInt(g, -12, -8, CLOTHES_DISTRIBUTION[:])
}
func getHandGene(g string) int {
	return getGeneInt(g, -16, -12, HAND_DISTRIBUTION[:])
}
func getHeadGene(g string) int {
	return getGeneInt(g, -20, -16, HEAD_DISTRIBUTION[:])
}
func getMouthGene(g string) int {
	return getGeneInt(g, -24, -20, MOUTH_DISTRIBUTION[:])
}
func getEyesGene(g string) int {
	return getGeneInt(g, -28, -24, EYES_DISTRIBUTION[:])
}

func getHeadGeneAttribute(g string, configService *config.ConfigService) StringAttribute {
	gene := getHeadGene(g)
	return StringAttribute{
		TraitType: "Head",
		Value:     configService.Head[gene],
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
		Value:     configService.Skin[gene],
	}
}
func getClothesGeneAttribute(g string, configService *config.ConfigService) StringAttribute {
	gene := getClothesGene(g)
	return StringAttribute{
		TraitType: "Clothes",
		Value:     configService.Clothes[gene],
	}
}
func getHandGeneAttribute(g string, configService *config.ConfigService) StringAttribute {
	gene := getHandGene(g)
	return StringAttribute{
		TraitType: "Hand",
		Value:     configService.Hand[gene],
	}
}
func getBackgroundGeneAttribute(g string, configService *config.ConfigService) StringAttribute {
	gene := getBackgroundGene(g)
	return StringAttribute{
		TraitType: "Background",
		Value:     configService.Background[gene],
	}
}

func getMouthGeneAttribute(g string, configService *config.ConfigService) StringAttribute {
	gene := getMouthGene(g)
	return StringAttribute{
		TraitType: "Mouth",
		Value:     configService.Mouth[gene],
	}
}

func getHeadGenePath(g string) string {
	gene := getHeadGene(g)
	return Gene(gene).toPath()
}

func getEyesGenePath(g string) string {
	gene := getEyesGene(g)
	return Gene(gene).toPath()
}

func getSkinGenePath(g string) string {
	gene := getSkinGene(g)
	return Gene(gene).toPath()
}

func getClothesGenePath(g string) string {
	gene := getClothesGene(g)
	return Gene(gene).toPath()
}

func getHandGenePath(g string) string {
	gene := getHandGene(g)
	return Gene(gene).toPath()
}

func getBackgroundGenePath(g string) string {
	gene := getBackgroundGene(g)
	return Gene(gene).toPath()
}

func getMouthGenePath(g string) string {
	gene := getMouthGene(g)
	return Gene(gene).toPath()
}

func (g *Genome) name(configService *config.ConfigService, tokenId string) string {
	gStr := string(*g)
	gene := getSkinGene(gStr)
	return fmt.Sprintf("%v Lobby Lobster #%v", configService.Skin[gene], tokenId)
}

func (g *Genome) description(configService *config.ConfigService, tokenId string) string {
	gStr := string(*g)
	gene := getSkinGene(gStr)
	return fmt.Sprintf("The %v Lobby Lobster #%v is a citizen of the Polymorph Universe and has a unique genetic code!", configService.Skin[gene], tokenId)
}

func (g *Genome) genes() []string {
	gStr := string(*g)

	res := make([]string, 0, GENES_COUNT)

	res = append(res, getBackgroundGenePath(gStr))
	res = append(res, getSkinGenePath(gStr))
	res = append(res, getClothesGenePath(gStr))
	res = append(res, getHandGenePath(gStr))
	res = append(res, getHeadGenePath(gStr))
	res = append(res, getEyesGenePath(gStr))
	res = append(res, getMouthGenePath(gStr))

	return res
}

func (g *Genome) attributes(configService *config.ConfigService) []interface{} {
	gStr := string(*g)

	res := []interface{}{}
	res = append(res, getMouthGeneAttribute(gStr, configService))
	res = append(res, getSkinGeneAttribute(gStr, configService))
	res = append(res, getHandGeneAttribute(gStr, configService))
	res = append(res, getClothesGeneAttribute(gStr, configService))
	res = append(res, getEyesGeneAttribute(gStr, configService))
	res = append(res, getHeadGeneAttribute(gStr, configService))
	res = append(res, getBackgroundGeneAttribute(gStr, configService))
	return res
}

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

	// b.WriteString(".jpg") // Finish with jpg extension

	geneUrl := b.String()

	imageExists := resourceExists(fmt.Sprintf("%s.jpg", geneUrl))
	if !imageExists {
		generateAndSaveImage(genes)
	}

	videoExists := resourceExists(fmt.Sprintf("%s.mp4", geneUrl))
	if !videoExists {
		generateAndSaveVideo(tokenId, genes)
	}

	m.Image = geneUrl
	return m
}

type Metadata struct {
	Description string      `json:"description"`
	Name        string      `json:"name"`
	Image       string      `json:"image"`
	Attributes  interface{} `json:"attributes"`
	ExternalUrl string      `json:"external_url"`
}

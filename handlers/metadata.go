package handlers

import (
	"math/big"
	"net/http"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/go-chi/render"
	"github.com/lobster-metadata/config"
	"github.com/lobster-metadata/contracts"
	"github.com/lobster-metadata/dlt"
	"github.com/lobster-metadata/metadata"
	log "github.com/sirupsen/logrus"
)

func HandleMetadataRequest(ethClient *dlt.EthereumClient, address string, configService *config.ConfigService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		instance, err := contracts.NewLobster(common.HexToAddress(address), ethClient.Client)
		if err != nil {
			render.Status(r, 500)
			render.JSON(w, r, err)
			log.Errorln(err)
			return
		}

		tokenId := r.URL.Query().Get("id")

		iTokenId, err := strconv.Atoi(tokenId)
		if err != nil {
			render.Status(r, 500)
			render.JSON(w, r, err)
			log.Errorln(err)
			return
		}

		genomeInt, err := instance.GeneOf(nil, big.NewInt(int64(iTokenId)))
		if err != nil {
			render.Status(r, 500)
			render.JSON(w, r, err)
			log.Errorln(err)
			return
		}
		g := metadata.Genome(genomeInt.String())
		if g != "0" {
			render.JSON(w, r, (&g).Metadata(tokenId, configService))
		} else {
			render.JSON(w, r, Error{
				Status:  404,
				Message: "No lobster found"},
			)
		}
	}
}

type Error struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

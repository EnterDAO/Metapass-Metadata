package handlers

import (
	"math/big"
	"net/http"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/go-chi/render"
	"github.com/metapass-metadata/config"
	"github.com/metapass-metadata/contracts"
	"github.com/metapass-metadata/dlt"
	"github.com/metapass-metadata/metadata"
	log "github.com/sirupsen/logrus"
)

func HandleMetadataRequest(ethClient *dlt.EthereumClient, address string, configService *config.ConfigService, uniqueConfigService *config.ConfigService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		instance, err := contracts.NewShardedMinds(common.HexToAddress(address), ethClient.Client)
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

		isUnique, uniqueIndex, err := instance.IsTokenUnique(nil, big.NewInt(int64(iTokenId)))
		if err != nil {
			render.Status(r, 500)
			render.JSON(w, r, err)
			log.Errorln(err)
			return
		}


		g := metadata.Genome(genomeInt.String())
		if g != "0" {
			if !isUnique {
				render.JSON(w, r, (&g).Metadata(tokenId, configService))
			} else {
				render.JSON(w, r, (&g).UniqueMetadata(tokenId, uniqueIndex.Int64(), configService, uniqueConfigService))
			}
		} else {
			render.JSON(w, r, Error{
				Status:  404,
				Message: "No Sharded Mind found"},
			)
		}
	}
}

type Error struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

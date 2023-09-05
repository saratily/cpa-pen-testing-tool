package server

import (
	"cpa-pen-testing-tool/internal/store"
	"fmt"
	"net/http"
	"os/exec"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-ping/ping"
	nmap "github.com/lair-framework/go-nmap"
	"github.com/likexian/whois"
)

func indexTools(ctx *gin.Context) {

	paramID := ctx.Param("id")
	paramType := ctx.Param("type")

	user, err := currentUser(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": InternalServerError})
		return
	}
	penetration, err := store.FetchPenetration(paramID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if user.ID != penetration.UserID {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Not authorized."})
		return
	}
	tools, err := store.FetchPenTools(penetration, paramType)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "Tools fetched successfully.",
		"data": tools,
	})
}

func executeTool(ctx *gin.Context) {
	paramID := ctx.Param("id")
	paramType := ctx.Param("type")

	user, err := currentUser(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": InternalServerError})
		return
	}
	penetration, err := store.FetchPenetration(paramID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if user.ID != penetration.UserID {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Not authorized."})
		return
	}

	tools, err := store.FetchPenTools(penetration, paramType)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	for i, tool := range tools {
		switch tool.Type {
		case "reachable":
			response, errors := http.Get(tool.Options)

			if errors != nil {
				_, netErrors := http.Get("https://www.google.com")

				if netErrors != nil {
					tools[i].Output = "Error: no internet\n"
				}
			}

			if response.StatusCode == 200 {
				tools[i].Output = "Site is reachable\n"
			} else {
				tools[i].Output = "Error: Site is not reachable\n"
			}

		case "whois":
			tools[i].Output, err = whois.Whois(tool.Options)
			if err != nil {
				tools[i].Output = err.Error()
			}
		case "ping":
			pinger, err := ping.NewPinger("www.google.com")
			if err != nil {
				panic(err)
			}
			pinger.Count = 3
			pinger.Run()                 // blocks until finished
			stats := pinger.Statistics() // get send/receive/rtt stats
			tools[i].Output = fmt.Sprintf("%d packets transmitted, %d received, %f% packet loss, time 4087ms \nrtt min/avg/max/mdev = %s/%s/%s/%s ms", stats.PacketsSent, stats.PacketsRecv, stats.PacketLoss, stats.MinRtt, stats.AvgRtt, stats.MaxRtt, stats.StdDevRtt)

		case "digIPv4":
			out, err := exec.Command("dig", tool.Options, "+short", "A").Output()
			if err != nil {
				tools[i].Output = err.Error()
			}
			tools[i].Output = fmt.Sprintf("%s", string(out))

		case "digIPv6":
			out, err := exec.Command("dig", tool.Options, "+short", "AAAA").Output()
			if err != nil {
				tools[i].Output = err.Error()
			}
			tools[i].Output = fmt.Sprintf("%s", string(out))

		case "digCNAME":
			out, err := exec.Command("dig", tool.Options, "CNAME").Output()
			if err != nil {
				tools[i].Output = err.Error()
			}
			tools[i].Output = fmt.Sprintf("%s", string(out))

		case "digMX":
			out, err := exec.Command("dig", tool.Options, "MX").Output()
			if err != nil {
				tools[i].Output = err.Error()
			}
			tools[i].Output = fmt.Sprintf("%s", string(out))

		case "digNS":
			out, err := exec.Command("dig", tool.Options, "NS").Output()
			if err != nil {
				tools[i].Output = err.Error()
			}
			tools[i].Output = fmt.Sprintf("%s", string(out))

		case "digTXT":
			out, err := exec.Command("dig", tool.Options, "TXT").Output()
			if err != nil {
				tools[i].Output = err.Error()
			}
			tools[i].Output = fmt.Sprintf("%s", string(out))

		case "digANY":
			out, err := exec.Command("dig", tool.Options, "ANY").Output()
			if err != nil {
				tools[i].Output = err.Error()
			}
			tools[i].Output = fmt.Sprintf("%s", string(out))

		case "digSOA":
			out, err := exec.Command("dig", tool.Options, "SOA").Output()
			if err != nil {
				tools[i].Output = err.Error()
			}
			tools[i].Output = fmt.Sprintf("%s", string(out))

		case "LookupIP":
			cmd := "nslookup " + tool.Options + " | awk '/^Address: / { print $2 }'"
			out, err := exec.Command("bash", "-c", cmd).Output()
			if err != nil {
				tools[i].Output = err.Error()
			}
			tools[i].Output = fmt.Sprintf("%s", string(out))
		case "LookupCNAME":
			out, err := exec.Command("nslookup", "-type=cname", tool.Options).Output()
			if err != nil {
				tools[i].Output = err.Error()
			}
			tools[i].Output = fmt.Sprintf("%s", string(out))
		case "LookupMX":
			out, err := exec.Command("nslookup", "-type=mx", tool.Options).Output()
			if err != nil {
				tools[i].Output = err.Error()
			}
			tools[i].Output = fmt.Sprintf("%s", string(out))
		case "LookupNS":
			out, err := exec.Command("nslookup", "-type=ns", tool.Options).Output()
			if err != nil {
				tools[i].Output = err.Error()
			}
			tools[i].Output = fmt.Sprintf("%s", string(out))
		case "LookupSRV":
			out, err := exec.Command("nslookup", "-type=srv", tool.Options).Output()
			if err != nil {
				tools[i].Output = err.Error()
			}
			tools[i].Output = fmt.Sprintf("%s", string(out))
		case "LookupTXT":
			out, err := exec.Command("nslookup", "-type=txt", tool.Options).Output()
			if err != nil {
				tools[i].Output = err.Error()
			}
			tools[i].Output = fmt.Sprintf("%s", string(out))

		case "nmap":
			content, err := exec.Command("nmap", "-sT", "google.com").Output()
			if err != nil {
				tools[i].Output = err.Error()
			}

			out, _ := nmap.Parse(content)
			fmt.Println(out)

			// tools[i].Output = fmt.Sprintf("%s", string(out))

		default:
			out, err := exec.Command(tool.Options).Output()
			if err != nil {
				tools[i].Output = err.Error()
			}
			tools[i].Output = fmt.Sprintf("%s", string(out))
		}

		tools[i].ModifiedAt = time.Now()
		tools[i].Selected = 1
		if err := store.UpdateTool(&tools[i]); err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": InternalServerError})
			return
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "Tools fetched successfully.",
		"data": tools,
	})
}

func toggleTool(ctx *gin.Context) {
	paramID := ctx.Param("id")
	paramType := ctx.Param("type")

	user, err := currentUser(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": InternalServerError})
		return
	}
	penetration, err := store.FetchPenetration(paramID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if user.ID != penetration.UserID {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Not authorized."})
		return
	}

	tools, err := store.FetchPenTools(penetration, paramType)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	for i, tool := range tools {
		if tool.Selected == 1 {
			tools[i].Selected = 2
		} else {
			tools[i].Selected = 1
		}
		if err := store.UpdateTool(&tools[i]); err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": InternalServerError})
			return
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "Tools fetched successfully.",
		"data": tools,
	})
}

/*
// func createTool(ctx *gin.Context) {
// 	tool := ctx.MustGet(gin.BindKey).(*store.Tool)
// 	user, err := currentUser(ctx)
// 	if err != nil {
// 		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}
// 	if err := store.AddTool(user, tool); err != nil {
// 		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	ctx.JSON(http.StatusOK, gin.H{
// 		"msg":  "Tool created successfully.",
// 		"data": tool,
// 	})
// }


func updateTool(ctx *gin.Context) {
	// jsonTool := ctx.MustGet(gin.BindKey).(*store.Tool)
	// user, err := currentUser(ctx)
	// if err != nil {
	// 	ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": InternalServerError})
	// 	return
	// }
	// dbTool, err := store.FetchTool(jsonTool.ID)
	// if err != nil {
	// 	ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }
	// if user.ID != dbTool.UserID {
	// 	ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Not authorized."})
	// 	return
	// }
	// jsonTool.ModifiedAt = time.Now()
	// if err := store.UpdateTool(jsonTool); err != nil {
	// 	ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": InternalServerError})
	// 	return
	// }
	// ctx.JSON(http.StatusOK, gin.H{
	// 	"msg":  "Tool updated successfully.",
	// 	"data": jsonTool,
	// })
}

func deleteTool(ctx *gin.Context) {
	// paramID := ctx.Param("id")
	// id, err := strconv.Atoi(paramID)
	// if err != nil {
	// 	ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Not valid ID."})
	// 	return
	// }
	// user, err := currentUser(ctx)
	// if err != nil {
	// 	ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": InternalServerError})
	// 	return
	// }
	// tool, err := store.FetchTool(id)
	// if err != nil {
	// 	ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }
	// if user.ID != tool.UserID {
	// 	ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Not authorized."})
	// 	return
	// }
	// if err := store.DeleteTool(tool); err != nil {
	// 	ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }
	ctx.JSON(http.StatusOK, gin.H{"msg": "Tool deleted successfully."})
}
*/

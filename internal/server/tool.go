package server

import (
	"cpa-pen-testing-tool/internal/store"
	"fmt"
	"net/http"
	"os/exec"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
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
		case "ping":
			out, err := exec.Command("ping", tool.Options, "-c 5", "-i 3", "-w 10").Output()
			if err != nil {
				tools[i].Output = err.Error()
			}
			if strings.Contains(string(out), "Destination Host Unreachable") {
				tools[i].Output = "Destination Host Unreachable"
			} else {
				tools[i].Output = "Destination Host is reachable"
			}
		case "whois":
			tools[i].Output, err = whois.Whois(tool.Options)
			if err != nil {
				tools[i].Output = err.Error()
			}
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
			out, err := exec.Command("nmap", "-sT", "google.com").Output()
			if err != nil {
				tools[i].Output = err.Error()
			}
			tools[i].Output = fmt.Sprintf("%s", string(out))

		case "nikto":
			out, err := exec.Command("nikto", "-h", tool.Options).Output()
			if err != nil {
				tools[i].Output = err.Error()
			}
			tools[i].Output = fmt.Sprintf("%s", string(out))

		case "ffuf":
			/*ffuf -u https://FUZZ.yahoo.com -w /usr/share/wordlists/dirb/common.txt -p 1 fc 301
			ffuf -u https://api.yahoo.com/FUZZ -w /usr/share/wordlists/dirb/common.txt -p 1
			*/
			out, err := exec.Command("ffuf", "-u", "http://"+tool.Options+"/FUZZ", "-w", "/usr/share/wordlists/dirb/common.txt", "-p", "1").Output()
			if err != nil {
				tools[i].Output = err.Error()
			}
			tools[i].Output = fmt.Sprintf("%s", string(out))
		case "dirb":
			fmt.Print(tool.Options)
			//dirb http://example.com -w /usr/share/wordlists/dirb/common.txt
			out, err := exec.Command("dirb", "http://"+tool.Options, "-w", "/usr/share/wordlists/dirb/common.txt").Output()
			if err != nil {
				tools[i].Output = err.Error()
			}
			tools[i].Output = fmt.Sprintf("%s", string(out))
		case "wfuzz":
			//wfuzz -c -w /usr/share/wordlists/dirb/common.txt https://example.com/FUZZ
			// wfuzz -c -w /usr/share/wordlists/dirb/common.txt -R 1 https://{FUZZ}.example.com/
			//usr/share/wordlists/wfuzz/wordlist.txt
			//wfuzz -w wordlist.txt -f output.txt --hc 404 --follow http://facebook.com/FUZZ
			// out, err := exec.Command("wfuzz", "-w", "usr/share/wordlists/wfuzz/wordlist.txt", "-hc", "404", "--follow", "http://"+tool.Options+"/FUZZ").Output()
			// out, err := exec.Command("wfuzz", "-z", "list,GET-HEAD-POST-TRACE-OPTIONS", "-X", "FUZZ", "http://"+tool.Options+"/").Output()
			//wfuzz -z list,GET-HEAD-POST-TRACE-OPTIONS -X FUZZ http://testphp.vulnweb.com/

			out, err := exec.Command("wfuzz", "-c", "-w", "/usr/share/wordlists/dirb/common.txt", "https://"+tool.Options+"/FUZZ").Output()

			if err != nil {
				tools[i].Output = err.Error()
			}
			tools[i].Output = fmt.Sprintf("%s", string(out))
		case "wappalyzer":
			out, err := exec.Command("wappy", "-u", tool.Options).Output()

			if err != nil {
				tools[i].Output = err.Error()
			}
			tools[i].Output = fmt.Sprintf("%s", string(out))

		default:
			tools[i].Output = "Unknown tool"
			// out, err := exec.Command(tool.Options).Output()
			// if err != nil {
			// 	tools[i].Output = err.Error()
			// }
			// tools[i].Output = fmt.Sprintf("%s", string(out))
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

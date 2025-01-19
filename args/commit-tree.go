package args

import (
	"bytes"
	"fmt"
	"os"
	"time"

	"github.com/isantoshgyawali/git-go/utils"
)

func CommitTree(message string, author *utils.Author) (string, error) {
  _, rootPath, err := utils.FindGitRoot()
  if err != nil {
    fmt.Fprintf(os.Stderr, err.Error())
  }

  treeHash, err := WriteTree(rootPath)
  if err != nil {
    return "", err 
  }

  var parentCommit string

  message = "initial-commit"
  unixTime := time.Now().Unix()
  timeZone := time.Now().Format("-07:00") // there exist a predefined format : "2006-01-02 15:04:05 -07:00"

  fmt.Println(unixTime, timeZone)

  // refrence commit:
  // tree a0579abd9807dec67b605c9f1f84949afd85d41a
  // parent cc6e107ad9b830ab03550595c53ad5d3401f5988
  // author Santosh Gyawali <isantoshgyawali@gmail.com> 1737281355 +0545
  // committer Santosh Gyawali <isantoshgyawali@gmail.com> 1737281355 +0545
  //
  // initial-commit

  commitContent := fmt.Sprintf(
    "tree %s\n%s"+
    "author %s <%s> %d %s\n"+
    "committer %s <%s> %d %s\n\n"+
    "%s\n", 
    treeHash,
    func() string{
      if parentCommit != "" {
        return fmt.Sprintf("parent %s\n", parentCommit)
      }
      return ""
    }(),
    author.Name, author.Email, unixTime, timeZone,
    author.Name, author.Email, unixTime, timeZone,
    message,
  )

  HeaderAndData := fmt.Sprintf("commit %d\000%s", len(commitContent), commitContent)
  hash := utils.HashIt([]byte(HeaderAndData))
  commitHash := fmt.Sprintf("%x", hash)

  return commitHash, nil
}

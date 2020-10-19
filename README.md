# GO_api_appointy

Packages Used:

import (
	"context"
	"encoding/json"
	"fmt"

	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/mongo/options"
)

Operations completed:

1) Nested Meetings and participants structure.
2) Post request to schedule meeting(Route: /meetings)
3) Get request to display all the meetings(Route: /meetingsall)

Parital IMplementation:

1) Get request to get meeting based on ID.

Get snippet on postman:






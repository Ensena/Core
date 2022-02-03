package core

import (
	"encoding/json"
	"fmt"

	bot "github.com/Ensena/bot-telegram"
	"github.com/Ensena/core/graphql-client"
	"github.com/gin-gonic/gin"
)

type userEnsena struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Picture   string `json:"picture"`
	Public    bool   `json:"public"`
	MoodleUDP bool   `json:"moodleUdp"`
	Cover     string `json:"cover"`
	AboutMe   string `json:"aboutMe"`
}

type UserMe struct {
	Data struct {
		User userEnsena `json:"user"`
	} `json:"data"`
}

func GetEnsenaData(ctx *gin.Context, userID int) ([]byte, userEnsena) {

	graphMe := fmt.Sprintf(`{
		allApps {
		  edges {
			node {
			  id
			  name
			  description
			  usersAppsByAppId(condition: { userId: %d ,enable:true }) {
				enable: totalCount
			   }
			  userByOwnerId {
				name
				lastName
				email
			  }
			}
		  }
		}
		user: userById(id: %d) {
		  id
		  name
		  lastName
		  email
		  picture
		  moodleUdp
		  public
		  cover
		  aboutMe
		  role
		  UDP: authsByUserId(condition: { system: 1 }) {
			edges {
			  node {
				custom
			  }
			}
		  }
	  
		  sectionsByOwnerId {
			edges {
			  node {
				id
				semester
				enable
				year
				section
				custom
				filesSectionsBySectionId {
				  edges {
					node {
					  fileByFileId {
						id
						name
						url
						userByOwnerId {
						  id
						  name
						  lastName
						}
					  }
					}
				  }
				}
				tasksBySectionId {
				  edges {
					node {
					  id
					  name
					  description
					  start
					  finish
					  integrates
					  role
					  taskFilesByTaskId {
						edges {
						  node {
							id
							fileByFileId {
							  id
							}
						  }
						}
					  }
					}
				  }
				}
				usersSectionsBySectionId {
				  edges {
					node {
					  id
					  role
					  userByUserId {
						id
						name
						lastName
						email
						picture
					  }
					}
				  }
				}
				messagesBySectionId(
				  condition: { forum: true }
				  first: 4
				  orderBy: ID_DESC
				) {
				  edges {
					node {
					  id
					  title
					  message
					  messageByMessageId {
						id
						title
						message
						userByOwnerId {
						  id
						  name
						  lastName
						}
					  }
					  userByOwnerId {
						id
						name
						lastName
					  }
					}
				  }
				}
				courseByCourseId {
				  id
				  name
				  code
				  institutionByOwnerInstitutionId {
					id
					name
				  }
				}
			  }
			}
		  }
	  
		  usersSectionsByUserId {
			edges {
			  node {
				role
				sectionBySectionId {
				  id
				  semester
				  year
				  enable

				  section
				  custom
				  tasksBySectionId(condition: { enable: true }) {
					edges {
					  node {
						id
						name
						description
						start
						finish
						integrates
						role
						taskFilesByTaskId {
						  edges {
							node {
							  id
							  fileByFileId {
								id
							  }
							}
						  }
						}
					  }
					}
				  }
				  oldTasks: tasksBySectionId(condition: { enable: false }) {
					edges {
					  node {
						id
						name
						description
						taskAnswersByTaskId(condition: { ownerId: %d }) {
						  edges {
							node {
							  id
							  review
							  calification
							}
						  }
						}
	  
						taskFilesByTaskId {
						  edges {
							node {
							  id
							  fileByFileId {
								id
							  }
							}
						  }
						}
					  }
					}
				  }
				  userByOwnerId {
					id
					name
					lastName
					email
					aboutMe
				  }
				  usersSectionsBySectionId {
					edges {
					  node {
						id
						role
						userByUserId {
						  id
						  name
						  lastName
						  email
						}
					  }
					}
				  }
				  messagesBySectionId(
					condition: { forum: true }
					first: 4
					orderBy: ID_DESC
				  ) {
					edges {
					  node {
						id
						title
						message
						messageByMessageId {
						  id
						  title
						  message
						  userByOwnerId {
							id
							name
							lastName
						  }
						}
						userByOwnerId {
						  id
						  name
						  lastName
						}
					  }
					}
				  }
	  
				  courseByCourseId {
					id
					name
					code
					description
					video
					institutionByOwnerInstitutionId {
					  id
					  name
					}
				  }
	  
				  filesSectionsBySectionId {
					edges {
					  node {
						fileByFileId {
						  id
						  name
						  url
						}
					  }
					}
				  }
				}
			  }
			}
		  }
		}
	  }
	  
	  `, userID, userID, userID)
	response, err := graphql.Query(ctx, graphMe)
	us := UserMe{}
	if err == nil {
		err = json.Unmarshal(response, &us)
		if err == nil {
			msg := fmt.Sprintf(`%s %s ha cargado Información Base
			Correo  :  %s`, us.Data.User.Name, us.Data.User.LastName, us.Data.User.Email)
			go bot.Send(bot.SendID, msg)
		}
	}
	return response, us.Data.User
}

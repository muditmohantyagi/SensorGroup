package controllers

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"test.com/config"
	"test.com/models"
	"test.com/pkg/helper"
)

type GroupController struct{}

var db = config.GoConnect()

// CreateGroups godocs
// @Summary Generate groups and sensors name API-1
// @Description It will save data of group and sensors name in DataBase
// @Produce  application/json
// @Tags Kicke-off phase
// @Router /group_sensor_generation [post]
func (con GroupController) GroupSensorGeneration(c *gin.Context) {
	var Sensors []models.Group

	Group := map[int]string{1: "alpha", 2: "beta", 3: "gamma"}
	for _, v := range Group {
		var Sensor models.Group
		max_group_data, is_data := models.MaxSensor(v)
		if !is_data {
			Sensor.GroupName = v
			Sensor.Code = 1
			Sensor.CodeName = v + " 1"

		} else {
			Sensor.GroupName = v
			Sensor.Code = max_group_data.Code + 1
			code_string := strconv.Itoa(Sensor.Code)
			Sensor.CodeName = v + " " + code_string

		}
		Sensors = append(Sensors, Sensor)

	}
	if result := db.Create(&Sensors); result.Error != nil {
		response := helper.Error("SQL Error1", result.Error.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, response)
		return
	}
	//c.Header("Content-Type", "application/json")
	response := helper.Success(true, "ok", "Groups and sensors created successfully")
	c.JSON(http.StatusOK, response)
}

// GenerateData godocs
// @Summary Generate data for groups and sensors API-2
// @Description It will save data  in DB for temparature, fish, cordinates etc
// @Produce  application/json
// @Tags Data generation face
// @Router /data_generation [post]
func (con GroupController) DataGeneration(c *gin.Context) {

	//var Sensors []models.Sensor

	Group := map[int]string{1: "alpha", 2: "beta", 3: "gamma"}
	SpeciesNames := map[int]string{1: "Atlantic Bluefin Tuna", 2: "Atlantic Cod", 3: "Atlantic Goliath Grouper", 4: "Atlantic Salmon", 5: "Atlantic Trumpetfish",
		6: "Atlantic Wolffish", 7: "Banded Butterflyfish", 8: "Beluga Sturgeon", 9: "Blue Marlin", 10: "Blue Tang", 11: "Bluebanded Goby", 12: "Bluehead Wrasse",
		13: "California Grunion", 14: "Chilean Common Hake", 15: "Chilean Jack Mackerel",
	}
	for _, v := range Group {
		randNum := rand.Intn(60)
		if randNum == 0 {
			randNum = 1
		}
		sensor_data, err := models.SensorByGroup(v)
		if err != nil {
			response := helper.Error("SQL Error1", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusBadRequest, response)
			return
		}
		// response2 := helper.Success(true, "ok", sensor_data)
		// c.JSON(http.StatusOK, response2)
		for _, s_data := range sensor_data {
			fmt.Println(s_data)
			randNum2 := rand.Intn(10)
			if randNum2 == 0 {
				randNum2 = 1
			}
			var Sensor models.Sensor

			//sensors
			randNum = randNum + randNum2

			Sensor.X = randNum + randNum2 + 1
			Sensor.Y = randNum + randNum2 + 2
			Sensor.Z = randNum + randNum2 + 3
			var Temp float32
			Sensor.Temperature = 1.1
			if Sensor.Z <= 10 {
				Temp = 1.1
			} else if Sensor.Z <= 20 {
				Temp = 2.3
			} else if Sensor.Z <= 30 {
				Temp = 8.2
			} else if Sensor.Z <= 40 {
				Temp = 10.2
			} else if Sensor.Z <= 50 {
				Temp = 12.2
			} else if Sensor.Z <= 60 {
				Temp = 15.2
			} else if Sensor.Z <= 70 {
				Temp = 17.2
			} else {
				Temp = 20
			}

			Sensor.Temperature = Temp
			Sensor.Transparency = int(Temp)
			Sensor.GroupID = s_data.ID
			if result := db.Create(&Sensor); result.Error != nil {
				response := helper.Error("SQL Error2", result.Error.Error(), helper.EmptyObj{})
				c.JSON(http.StatusBadRequest, response)
				return
			}
			//fishes
			var species_data []models.Species
			Range := rand.Intn(4)
			if Range == 0 {
				Range = 1
			}
			for i := 1; i <= Range; i++ {
				var Species models.Species
				randNum4 := rand.Intn(15)
				if randNum4 == 0 {
					randNum4 = 1
				}
				randNumFish := randNum4
				Species.Name = SpeciesNames[randNumFish]
				Species.Count = randNumFish
				Species.SensorID = Sensor.ID
				species_data = append(species_data, Species)
			}
			if result := db.Create(&species_data); result.Error != nil {
				response := helper.Error("SQL Error3", result.Error.Error(), helper.EmptyObj{})
				c.JSON(http.StatusBadRequest, response)
				return
			}

		}

	}
	response := helper.Success(true, "ok", "Data generated sussessfully")
	c.JSON(http.StatusOK, response)
}

// ShowFullData godocs
// @Summary show full data available in DB API-General
// @Description It will display whole data in database
// @Produce  application/json
// @Tags Aggregate statistics
// @Router /data_list [post]
func (con GroupController) ShowAllData(c *gin.Context) {
	var Group []models.Group
	if result := db.Preload("Sensor").Preload("Sensor.Species").Order("group_name").Find(&Group); result.Error != nil {
		response := helper.Error("SQL Error", result.Error.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.Success(true, "ok", Group)
	c.JSON(http.StatusOK, response)
}

///group/<groupName>/transparency/average:
// ShowAverageTransparency godocs
// @Summary show average transparency in a group API-3.1
// @Description It will display transparency of a groups
// @Produce  application/json
// @param groupName path string true "generate avg transparency by group name"
// @Tags Aggregate statistics
// @Router /group/transparency/average/{groupName} [post]
func (con GroupController) Transparency(c *gin.Context) {
	//db = config.GoConnect().Debug()
	group := c.Param("groupName")
	if group == "" {
		response := helper.Error("URL Error", "inavlid param", helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, response)
		return
	}
	var Group models.Group

	if result := db.Where("group_name=?", group).Joins("LEFT JOIN sensors ON sensors.group_id = groups.id").Select("AVG(sensors.transparency) as comman_var").Take(&Group); result.Error != nil {
		response := helper.Error("SQL Error", result.Error.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.Success(true, "ok", Group.CommanVar)
	c.JSON(http.StatusOK, response)
}

// /group/<groupName>/temperature/average
// ShowAverageTemperature godocs
// @Summary show average temperature in a group API-3.1
// @Description It will display temperature of a groups
// @Produce  application/json
// @param groupName path string true "generate avg temperature by group name"
// @Tags Aggregate statistics
// @Router /group/temperature/average/{groupName} [post]
func (con GroupController) Temperature(c *gin.Context) {
	group := c.Param("groupName")
	if group == "" {
		response := helper.Error("URL Error", "inavlid param", helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, response)
		return
	}
	var Group models.Group

	if result := db.Where("group_name=?", group).Joins("LEFT JOIN sensors ON sensors.group_id = groups.id").Select("AVG(sensors.temperature) as comman_var").Take(&Group); result.Error != nil {
		response := helper.Error("SQL Error", result.Error.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.Success(true, "ok", Group.CommanVar)
	c.JSON(http.StatusOK, response)
}

//sensor/average/temperature/<codeName>?from=<fromDateTime>&till=<untillDateTime>
// @Summary Show average temperature in a CodeName, e.g., alpha 1
// @Description Display the average temperature of a CodeName within a specified duration
// @Produce application/json
// @Param codeName path string true "sensor CodeName"
// @Param from query integer true "Unix timestamp for from of duration"
// @Param till query integer true "Unix timestamp for till of duration"
// @Tags Aggregate statistics
// @Router /sensor/average/temperature/{codeName} [post]
func (con GroupController) SenserTemperature(c *gin.Context) {
	codeName := c.Param("codeName")
	from := c.Query("from")
	till := c.Query("till")
	if codeName == "" && from == "" && till == "" {
		response := helper.Error("URL Error", "inavlid param", helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, response)
		return
	}
	var Group models.Group

	if result := db.Where("code_name=? and sensors.updated>=? and sensors.updated<=?", codeName, from, till).Joins("LEFT JOIN sensors ON sensors.group_id = groups.id").Select("AVG(sensors.temperature) as comman_var").Take(&Group); result.Error != nil {
		response := helper.Error("SQL Error", result.Error.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.Success(true, "ok", Group.CommanVar)
	c.JSON(http.StatusOK, response)
}

///group/species/<groupName>GroupSpecies
// @Summary shows all species in agroup ex:alpha
// @Description shows all species in agroup that is name and count of fishes
// @Produce application/json
// @Param groupName path string true "Group name as alpha"
// @Tags Aggregate statistics
// @Router /group/species/{groupName} [post]
func (con GroupController) GroupSpecies(c *gin.Context) {
	group := c.Param("groupName")
	if group == "" {
		response := helper.Error("URL Error", "inavlid param", helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, response)
		return
	}
	var Group []models.Group
	if result := db.Where("group_name=?", group).Preload("Sensor").Preload("Sensor.Species").Order("group_name").Find(&Group); result.Error != nil {
		response := helper.Error("SQL Error", result.Error.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.Success(true, "ok", Group)
	c.JSON(http.StatusOK, response)
}

//group/species/top/<groupName>/<N>TopSpecies
// @Summary show top count of fish in a group optionaly supports duration
// @Description show top count of fish in a group
// @Produce application/json
// @Param codeName path string true "Group Name"
// @Param N path integer true "Max Count of fish you want to see"
// @Param from query integer false "Unix timestamp for from of duration"
// @Param till query integer false "Unix timestamp for till of duration"
// @Tags Aggregate statistics
// @Router /group/species/top/{codeName}/{N} [post]
// Ex:http://localhost:8080/api/group/species/top/alpha/12?from=1699773927&till=1699777728
func (con GroupController) TopSpecies(c *gin.Context) {
	group := c.Param("groupName")
	N := c.Param("N")
	from := c.Query("from")
	till := c.Query("till")
	if group == "" && N == "" {
		response := helper.Error("URL Error", "inavlid param", helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, response)
		return
	}
	var option bool
	if group != "" && N != "" && from != "" && till != "" {
		option = true
	}
	var id []int
	if result := db.Model(&models.Group{}).
		Joins("LEFT JOIN sensors ON sensors.group_id = groups.id").
		Where("groups.group_name=?", group).
		Preload("Sensor").
		Pluck("sensors.id", &id); result.Error != nil {
		response := helper.Error("SQL Error", result.Error.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, response)
		return
	}
	var Species []models.Species
	dbQ := db.Where("sensor_id IN ? and count>=?", id, N)
	if option {
		dbQ.Where("updated>=? and updated<=?", from, till)
	}
	if result := dbQ.Find(&Species); result.Error != nil {
		response := helper.Error("SQL Error", result.Error.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.Success(true, "ok", Species)
	c.JSON(http.StatusOK, response)
}

//region/temperature/min?xMin=<xMin>&xMax=<xMax>&yMin=<yMin>&yMax=<yMax>&zMin=<zMin>&zMax=<zMax>
//Ex: http://localhost:8080/api/region/temperature/min?xMin=10&xMax=100&yMin=10&yMax=100&zMin=10&zMax=100
// @Summary show minimum temprature in a reason
// @Description show minimum temprature in a min and max:x,y,z  reason
// @Produce application/json
// @Param xMin query integer true " x min cordinate"
// @Param xMax query integer true " x max cordinate"
// @Param yMin query integer true " y min cordinate"
// @Param yMax query integer true " y max cordinate"
// @Param zMin query integer true " z min cordinate"
// @Param zMax query integer true " z max cordinate"
// @Tags Aggregate statistics
// @Router /region/temperature/min [post]
func (con GroupController) MinTempReason(c *gin.Context) {

	xMin := c.Query("xMin")
	xMax := c.Query("xMax")
	yMin := c.Query("yMin")
	yMax := c.Query("yMax")
	zMin := c.Query("zMin")
	zMax := c.Query("zMax")
	if xMin == "" && xMax == "" && yMin == "" && yMax == "" && zMin == "" && zMax == "" {
		response := helper.Error("URL Error", "inavlid param", helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var db = config.GoConnect().Debug()
	var f float32
	if result := db.Model(&models.Sensor{}).Select("MIN(temperature)").Where("x between ? and ? and y between ? and ? and z between ? and ?", xMin, xMax, yMin, yMax, zMin, zMax).Take(&f); result.Error != nil {
		response := helper.Error("SQL Error", result.Error.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.Success(true, "ok", f)
	c.JSON(http.StatusOK, response)
}

//region/temperature/max?xMin=<xMin>&xMax=<xMax>&yMin=<yMin>&yMax=<yMax>&zMin=<zMin>&zMax=<zMax>func (con GroupController) MinTempReason(c *gin.Context) {
// @Summary show maximum temprature in a reason
// @Description show maximum temprature in a min and max:x,y,z reason
// @Produce application/json
// @Param xMin query integer true " x min cordinate"
// @Param xMax query integer true " x max cordinate"
// @Param yMin query integer true " y min cordinate"
// @Param yMax query integer true " y max cordinate"
// @Param zMin query integer true " z min cordinate"
// @Param zMax query integer true " z max cordinate"
// @Tags Aggregate statistics
// @Router /region/temperature/max [post]
func (con GroupController) MaxTempReason(c *gin.Context) {
	xMin := c.Query("xMin")
	xMax := c.Query("xMax")
	yMin := c.Query("yMin")
	yMax := c.Query("yMax")
	zMin := c.Query("zMin")
	zMax := c.Query("zMax")
	if xMin == "" && xMax == "" && yMin == "" && yMax == "" && zMin == "" && zMax == "" {
		response := helper.Error("URL Error", "inavlid param", helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, response)
		return
	}
	var db = config.GoConnect().Debug()
	var f float32
	if result := db.Model(&models.Sensor{}).Select("MAX(temperature)").Where("x between ? and ? and y between ? and ? and z between ? and ?", xMin, xMax, yMin, yMax, zMin, zMax).Take(&f); result.Error != nil {
		response := helper.Error("SQL Error", result.Error.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.Success(true, "ok", f)
	c.JSON(http.StatusOK, response)
}

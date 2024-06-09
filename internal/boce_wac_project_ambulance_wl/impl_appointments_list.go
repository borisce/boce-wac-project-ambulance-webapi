package boce_wac_project_ambulance_wl

import (
  "net/http"

  "github.com/gin-gonic/gin"
  "github.com/google/uuid"
  "github.com/borisce/boce-wac-project-ambulance-webapi/internal/db_service"
)


 // CreateAppointment - Creates a new appointment
 func (this *implAppointmentsListAPI) CreateAppointment(ctx *gin.Context) {
	value, exists := ctx.Get("db_service")
    if !exists {
        ctx.JSON(
            http.StatusInternalServerError,
            gin.H{
                "status":  "Internal Server Error",
                "message": "db not found",
                "error":   "db not found",
            })
        return
    }

    db, ok := value.(db_service.DbService[AppointmentsList])
    if !ok {
        ctx.JSON(
            http.StatusInternalServerError,
            gin.H{
                "status":  "Internal Server Error",
                "message": "db context is not of required type",
                "error":   "cannot cast db context to db_service.DbService",
            })
        return
    }

    appointmentslistitem := AppointmentsList{}
    err := ctx.BindJSON(&appointmentslistitem)
    if err != nil {
        ctx.JSON(
            http.StatusBadRequest,
            gin.H{
                "status":  "Bad Request",
                "message": "Invalid request body",
                "error":   err.Error(),
            })
        return
    }

    if appointmentslistitem.Id == "" {
        appointmentslistitem.Id = uuid.New().String()
    }

    err = db.CreateDocument(ctx, appointmentslistitem.Id, &appointmentslistitem)

    switch err {
    case nil:
        ctx.JSON(
            http.StatusCreated,
            appointmentslistitem,
        )
    case db_service.ErrConflict:
        ctx.JSON(
            http.StatusConflict,
            gin.H{
                "status":  "Conflict",
                "message": "Appoitment already exists",
                "error":   err.Error(),
            },
        )
    default:
        ctx.JSON(
            http.StatusBadGateway,
            gin.H{
                "status":  "Bad Gateway",
                "message": "Failed to create appoitment in database",
                "error":   err.Error(),
            },
        )
    }
 }

 // DeleteAppointment - Deletes an appointment
 func (this *implAppointmentsListAPI) DeleteAppointment(ctx *gin.Context) {
	value, exists := ctx.Get("db_service")
    if !exists {
        ctx.JSON(
            http.StatusInternalServerError,
            gin.H{
                "status":  "Internal Server Error",
                "message": "db_service not found",
                "error":   "db_service not found",
            })
        return
    }

    db, ok := value.(db_service.DbService[AppointmentsList])
    if !ok {
        ctx.JSON(
            http.StatusInternalServerError,
            gin.H{
                "status":  "Internal Server Error",
                "message": "db_service context is not of type db_service.DbService",
                "error":   "cannot cast db_service context to db_service.DbService",
            })
        return
    }

    id := ctx.Query("id")
    err := db.DeleteDocument(ctx, id)

    switch err {
    case nil:
		ctx.JSON(
			http.StatusOK,
			gin.H{
				"status":  "OK",
				"message": "Appointment deleted",
			},
		)
    case db_service.ErrNotFound:
        ctx.JSON(
            http.StatusNotFound,
            gin.H{
                "status":  "Not Found",
                "message": "Appoitment not found",
                "error":   err.Error(),
            },
        )
    default:
        ctx.JSON(
            http.StatusBadGateway,
            gin.H{
                "status":  "Bad Gateway",
                "message": "Failed to delete appoitment from database",
                "error":   err.Error(),
            })
    }
 }

 // GetAppointmentsList - Provides list of all appointments
 func (this *implAppointmentsListAPI) GetAppointmentsList(ctx *gin.Context) {
	value, exists := ctx.Get("db_service")
	if !exists {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "db not found",
				"error":   "db not found",
			})
		return
	}

	db, ok := value.(db_service.DbService[AppointmentsList])

	if !ok {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "db context is not of required type",
				"error":   "cannot cast db context to db_service.DbService",
			})
		return
	}

	appoitments, err := db.GetAllAppointments(ctx)

	if err != nil {
		ctx.JSON(
			http.StatusBadGateway,
			gin.H{
				"status":  "Bad Gateway",
				"message": "Cannot get list of appoitments",
				"error":   err.Error(),
			})
		return
	}

	ctx.JSON(
		http.StatusOK,
		appoitments,
	)
 }
 

 // UpdateAppointment - Updates an appointment
 func (this *implAppointmentsListAPI) UpdateAppointment(ctx *gin.Context) {
	value, exists := ctx.Get("db_service")
	if !exists {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "db not found",
				"error":   "db not found",
			})
		return
	}

	db, ok := value.(db_service.DbService[AppointmentsList])

	if !ok {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "db context is not of required type",
				"error":   "cannot cast db context to db_service.DbService",
			})
		return
	}

	appointment := AppointmentsList{}
	err := ctx.BindJSON(&appointment)

	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":  "Bad Request",
				"message": "Invalid request body",
				"error":   err.Error(),
			})
		return
	}

	err = db.UpdateDocument(ctx, appointment.Id, &appointment)

	switch err {
	case nil:
		ctx.JSON(
			http.StatusOK,
			appointment,
		)
	case db_service.ErrNotFound:
		ctx.JSON(
			http.StatusNotFound,
			gin.H{
				"status":  "Not Found",
				"message": "Document not found",
				"error":   err.Error(),
			},
		)
	default:
		ctx.JSON(
			http.StatusBadGateway,
			gin.H{
				"status":  "Bad Gateway",
				"message": "Cannot update appointment",
				"error":   err.Error(),
			},
		)
	}


 }

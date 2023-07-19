function getReservationDates(args) {
    let {CSRFToken = ""} = args
    const formId = "reservation-dates-form"
    let html = `<form class="reservation-dates-form" id="reservation-dates-form" action="reservation.html" method="GET" class="needs-validation" novalidate>
          <div class= "row g-3 p-3" >

      <div class="col">
        <label for="start" class="form-label">Starting Date</label>
        <input disabled required type="text" class="form-control" name="start" id="start" placeholder="Arrival">
      </div>

      <div class="col">
        <label for="end" class="form-label">Ending Date</label>
        <input disabled required type="text" class="form-control" name="end" id="end" placeholder="Departure">
      </div>
      </div>
      </form > `
  
     Prompt().reservationDates({
      message: html,
      title: "Choose Your Dates",
      confirmButtonText: "Check Availability",
      callback: async function (result) {
        const form = document.getElementById(formId)
        const formData = new FormData(form)
        formData.append("csrf_token", CSRFToken)

        const response = await fetch("/search-availability-json", {
          method: "post",
          body: formData
        })

        const data = response.json()
        console.log(data)
      },
      preConfirm: () => {
        return [
          document.getElementById('start').value,
          document.getElementById('end').value
        ]

      },
      didRender: function () {
        let datepickerForm = document.getElementById(formId)

        const datepicker = new DateRangePicker(datepickerForm, {
          format: "dd/mm/yyyy",
          container: ".swal2-popup"
        })
      },
      didOpen: function () {
        inputs = document.querySelectorAll("input")

        inputs.forEach(input => {
          input.disabled = false
        })
      }
    })

  }

function Prompt() {
    let toast = function (args) {
        let {
            title = "",
            message = "",
            icon = "success",
            position = 'top-end',
            showConfirmButton = false,
            timer = 3000,
            timerProgressBar = true,
        } = args

        const Toast = Swal.mixin({
            toast: true,
            position: position,
            showConfirmButton: showConfirmButton,
            timer: timer,
            timerProgressBar: timerProgressBar,
            didOpen: (toast) => {
                toast.addEventListener('mouseenter', Swal.stopTimer)
                toast.addEventListener('mouseleave', Swal.resumeTimer)
            }
        })

        Toast.fire({
            icon: icon,
            title: title,
            html: message,
        })
    }

    let alert = function (args) {
        let {
            icon = "",
            title = "",
            message = "",
            footer = ""
        } = args

        Swal.fire({
            icon: icon,
            title: title,
            html: message,
            footer: footer
        })
    }

    let success = function (args) {
        args["icon"] = "success"
        alert(args)
    }

    let error = function (args) {
        args["icon"] = "error"
        alert(args)
    }

    let reservationDates = async function (args) {
        let {
            message = "",
            title = "",
            confirmButtonText = "",
            didRender = undefined,
            preConfirm = undefined,
            didOpen = undefined,
            callback = undefined,
        } = args

        const { value: formValues } = await Swal.fire({
            title: title,
            html: message,
            focusConfirm: false,
            backdrop: false,
            confirmButtonText: confirmButtonText,
            didRender: () => {
                return didRender()
            },
            preConfirm: () => {
                return preConfirm()
            },
            didOpen: () => {
                return didOpen()
            }
        })

        if (callback == undefined) {
          return
        }

        if (formValues.dismiss === Swal.DismissReason.cancel) {
          callback(false)
          return
        }

        if (formValues.value == "") {
          callback(false)
          return
        }

        callback(true)
    }

    return {
        toast: toast,
        success: success,
        error: error,
        reservationDates: reservationDates
    }
}
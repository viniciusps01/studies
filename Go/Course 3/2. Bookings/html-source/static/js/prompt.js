function Prompt() {
    let toast = function (args) {
        let {
            title = "",
            message = "",
            icon = "success",
            position = 'top-end',
            showConfirmButton = false,
            timer = 3000,
            timerProgressBar = true
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

        if (formValues) {
            Swal.fire(JSON.stringify(formValues))
        }
    }

    return {
        toast: toast,
        success: success,
        error: error,
        reservationDates: reservationDates
    }
}
import Swal from 'sweetalert2'

export default {
  error(title, text) {
    Swal.fire({
      icon: 'error',
      title: title,
      text: text
    })
  },
  success(title, text) {
    Swal.fire({
      icon: 'success',
      title: title,
      text: text
    })
  },
  alert(text) {
    Swal.fire(text)
  },
  confirm() {
    Swal.fire({
      title: 'Are you sure?',
      text: "You won't be able to revert this!",
      icon: 'warning',
      showCancelButton: true,
      confirmButtonColor: '#3085d6',
      cancelButtonColor: '#d33',
      confirmButtonText: 'Yes, delete it!'
    }).then((result) => {
      if (result.value) {
        Swal.fire(
          'Deleted!',
          'Your file has been deleted.',
          'success'
        )
      }
    })
  }
}
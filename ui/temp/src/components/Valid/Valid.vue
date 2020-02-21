<template>
  <div>
    <b-form-invalid-feedback :state="false">
      Your user ID must be 5-12 characters long.
    </b-form-invalid-feedback>
    <b-form-valid-feedback :state="true">
      Looks Good.
    </b-form-valid-feedback>
  </div>
</template>
<template>
  <TBlock title="checkbox" subtitle="example">
    <p>
      <TButton @click="change1">change checkbox1</TButton>
      <TCheckbox v-model="checked1" @change="changeHandler" val="{}" name="checkbox1" label="example" />
    </p>
  </TBlock>
</template>
<script>
export default {
  data() {
    return {
      checked1: false
    }
  },
  methods: {
    change1() {
      this.checked1 = !this.checked1
    },
    changeHandler(e) {
      console.log(e)
    }
  }
}
</script>
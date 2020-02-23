<template>
  <div class="t-upload-box" @click="chooseFile">
    <TIcon v-if="!imgSrc" icon="plus" />
    <img v-if="imgSrc" :src="imgSrc" />
    <input type="file" class="d-none" ref="hiddenFile" @change="fileHandler" />
  </div>
</template>
<script>
import { uploadFile } from '@/api/upload'
export default {
  props: {
    image: {
      type: String,
      default: ''
    }
  },
  data() {
    return {
      imgSrc: null,
      url: null
    }
  },
  mounted() {
    this.imgSrc = JSON.parse(JSON.stringify(this.image))
  },
  methods: {
    chooseFile() {
      this.$refs.hiddenFile.click()
    },
    fileHandler() {
      const file = this.$refs.hiddenFile.files[0]
      uploadFile(file).then(res => {
        this.url = res.path
        this.imgSrc = this.assetsURL(res.path)
        this.$emit('uploaded', this.url)
      })
    }
  }
}
</script>
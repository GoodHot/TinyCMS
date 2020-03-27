<template>
  <div :class="classes" @click="chooseFile">
    <t-icon v-if="!imgSrc" icon="plus" />
    <img v-if="imgSrc" :src="assetsURL(imgSrc)" />
    <input type="file" class="d-none" ref="hiddenFile" @change="fileHandler" />
  </div>
</template>
<script>
import { uploadFile } from '@/api/upload'

export default {
  model: {
    prop: 'image',
    event: 'change.image'
  },
  props: {
    image: {
      type: String,
      default: ''
    },
    state: {
      type: [Boolean, Object],
      default: null
    }
  },
  data() {
    return {
      imgSrc: null,
      url: null
    }
  },
  computed: {
    classes () {
      return [
        't-upload-box',
        {
          't-upload-box-err': !this.state && this.state !== null
        }
      ]
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
  },
  watch: {
    image () {
      this.imgSrc = JSON.parse(JSON.stringify(this.image))
    }
  }
}
</script>
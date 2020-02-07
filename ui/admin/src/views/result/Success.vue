<template>
  <a-card :bordered="false" style="margin: -24px -24px 0px;">
    <result type="success" :title="title">
    </result>
  </a-card>
</template>

<script>
import { Result } from '@/components'
import { mixinDevice } from '@/utils/mixin.js'

const directionType = {
  horizontal: 'horizontal',
  vertical: 'vertical'
}

export default {
  name: 'Success',
  components: {
    Result
  },
  mixins: [mixinDevice],
  data () {
    return {
      title: '提交成功',
      directionType
    }
  }
}
</script>

<style scoped>

</style>

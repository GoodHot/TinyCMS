<template>
<vue-nestable v-model="nestableItems" @input="inputHandler" @change="changeHandler">
  <vue-nestable-handle slot-scope="{ item }" :item="item">
    <b-button class="expand">+</b-button>
    <t-icon icon="folder" class="mr-1" /> {{ item.text }}
  </vue-nestable-handle>
</vue-nestable>
</template>
<script>
import { VueNestable, VueNestableHandle } from 'vue-nestable'
import '@/assets/css/nestable.scss'

export default {
  components: {
    VueNestable,
    VueNestableHandle
  },
  data () {
    return {
      nestableItems: [
        {
          id: 0,
          text: 'Andy'
        }, {
          id: 1,
          text: 'Harry',
          children: [{
            id: 2,
            text: 'David'
          }]
        }, {
          id: 3,
          text: 'Lisa'
        }
      ]
    }
  },
  methods: {
    inputHandler (val) {
      console.log(val)
    },
    changeHandler (val, opt) {
      console.log(val, opt)
    }
  }
}
</script>
<template>
  <div>
    <t-block title="Nestable" theme>
      <t-nestable v-model="nestableItems" class="push">
        <template v-slot:default="props">
          <t-icon icon="folder" class="mr-1" /> {{ props.item.text }}
          <span class="float-right">ddd</span>
        </template>
      </t-nestable>
    </t-block>
  </div>
</template>
<script>
export default {
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
  }
}
</script>

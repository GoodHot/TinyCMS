<template>
  <div>
    <TButton @click="showSimple">simple</TButton>
    <TButton @click="simple = true; size='sm'">simall</TButton>
    <TButton @click="simple = true; size='lg'">large</TButton>
    <TButton @click="simple = true; size='xl'">ext large</TButton>
    <TButton @click="simple = true; size='xl'; position='top'">top</TButton>
    <TButton @click="simple = true; size='xl'; position='center'">center</TButton>
    <TModal :show="simple" title="一个简单的modal" @onclose="simple = false" :size="size" :position="position">
      <p>simple modal</p>
    </TModal>
  </div>
</template>
<script>
export default {
  data() {
    return {
      simple: false,
      size: null,
      position: null,
      animation: ['fade-in','pop-in','pop-out','slide-up','slide-right','slide-left','from-right','from-left']
    }
  },
  methods: {
    showSimple() {
      this.simple = true;
      // setTimeout(() => {
      //   this.simple = false;
      // }, 2000);
    }
  }
};
</script>
<template>
  <div>
    <h1>elements</h1>
    <h2>theme</h2>
    <TInput placeholder="default" />
    <br>
    <TInput placeholder="alt" theme="alt" />
    <h2>readonly</h2>
    <TInput placeholder="alt" readonly />
    <br>
    <TInput placeholder="alt" readonly theme="alt" />
    <h2>disabled</h2>
    <TInput placeholder="alt" disabled />
    <br>
    <TInput placeholder="alt" disabled theme="alt" />
    <h2>state</h2>
    <TInput placeholder="success" state="success" />
    <br>
    <TInput placeholder="success" state="error" />
    <br>
    <TInput placeholder="success" state="success" theme="alt" />
    <br>
    <TInput placeholder="success" state="error" theme="alt" />
    <h2>textarea</h2>
    <TTextarea>hello</TTextarea><br>
    <TTextarea theme="alt">hello</TTextarea><br>
    <TTextarea state="success">hello</TTextarea><br>
    <TTextarea state="error">hello</TTextarea><br>
    <TTextarea theme="alt" state="success">hello</TTextarea><br>
    <TTextarea theme="alt" state="error">hello</TTextarea>
    <h2>switch</h2>
    <TSwitch v-for="item of types" :key="item" :type="item" :name="item" :label="item" /><br>
    <TSwitch v-for="item of types" :key="item" :type="item" :name="item" :label="item" size="lg" /><br>
    <TSwitch v-for="item of types" :key="item" :type="item" :name="item" :label="item" size="sm" />
  </div>
</template>
<script>
export default {
  data() {
    return {
      types: ['primary','secondary','success','info','warning','danger','dark','light']
    }
  }
}
</script>
<template>
  <div>
    <h1>Style</h1>
    <div class="row">
        <div class="col-md-6 col-xl-3">
            <TBlock title="TITLE" subtitle="Subtitle" ><p>Simple block..</p></TBlock>
        </div>
        <div class="col-md-6 col-xl-3">
            <TBlock title="TITLE" subtitle="Subtitle" headerBg ><p>With header background..</p></TBlock>
        </div>
        <div class="col-md-6 col-xl-3">
            <TBlock title="TITLE" subtitle="Subtitle" border><p>Bordered block..</p></TBlock>
        </div>
        <div class="col-md-6 col-xl-3">
            <TBlock title="TITLE" subtitle="Subtitle" headerBg border><p>Bordered block with header background..</p></TBlock>
        </div>
        <div class="col-md-6 col-xl-3">
            <TBlock title="TITLE" subtitle="Subtitle" borderRound><p>Border Round block..</p></TBlock>
        </div>
        <div class="col-md-6 col-xl-3">
            <TBlock title="TITLE" subtitle="Subtitle" borderRound shadow><p>Shadow block..</p></TBlock>
        </div>
        <div class="col-md-6 col-xl-3">
            <TBlock title="TITLE" subtitle="Subtitle" :transparent="true">Transparent</TBlock>
        </div>
    </div>
    <h1>Options</h1>
    <div class="row">
      <div class="col-md-6 col-xl-3">
        <TBlock title="TITLE" headerBg><p>Only title</p></TBlock>
      </div>
      <div class="col-md-6 col-xl-3">
        <TBlock subtitle="Subtitle" headerBg><p>Only subtitle</p></TBlock>
      </div>
      <div class="col-md-6 col-xl-3">
        <TBlock title="标题" subtitle="子标题" headerBg>
          <template slot="options">
            <TButton icon="plus" />
            <TButton iconPack="si" icon="fire" />
            <TButton iconPack="si" icon="ban" />
          </template>
          <p>Options</p>
        </TBlock>
      </div>
    </div>
    <h1>Themes</h1>
    <div class="row">
      <div class="col-md-6 col-xl-3" v-for="item of themes" :key="item">
        <TBlock :theme="item" title="title" ><p>{{ item }}</p></TBlock>
      </div>
    </div>
  </div>
</template>
<script>
export default {
  data() {
    return {
      themes: ['primary','primary-light','primary-dark','primary-darker','success','info','warning','danger','gray','gray-dark','gray-darker','black','default','default-light','default-dark','default-darker','amethyst','amethyst-light','amethyst-dark','amethyst-darker','city','city-light','city-dark','city-darker','flat','flat-light','flat-dark','flat-darker','modern','modern-light','modern-dark','modern-darker','smooth','smooth-light','smooth-dark','smooth-darker']
    }
  }
}
</script>
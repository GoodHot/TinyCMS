<template>
  <div class="row">
    <div class="col-4 push">
      <div class="input-group">
          <input type="text" class="form-control form-control-sm" placeholder="用户名搜索">
          <div class="input-group-append">
              <span class="input-group-text">
                  <TIcon icon="search" />
              </span>
          </div>
      </div>
    </div>
    <div class="col-8 push">
      <b-button variant="success" size="sm"><TIcon icon="plus" /> 创建新用户</b-button>
      <b-button variant="secondary" class="ml-1" size="sm" @click="$router.push('/role')"><TIcon icon="user" pack="si" /> 角色管理</b-button>
    </div>
    <div class="col-sm-6 col-xl-4">
      <UserCard>是打发斯蒂芬</UserCard>
    </div>
    <div class="col-sm-6 col-xl-4">
      <UserCard>不同渠道买到三盒3m n95，寄给国内亲戚。三个亲戚一人一盒。两盒8210外包装精美，一盒8210plus印刷粗糙。现在不知道怎么办了，真怕中招。平均一盒1000，邮费1000，破财了。一般上了货就是疯抢，不到2天就没了。</UserCard>
    </div>
    <div class="col-sm-6 col-xl-4">
      <UserCard>这次的疫情，城市像得了斯德哥尔综合症：一面被病毒所困呻吟怪叫，一面又浸淫在因此带来那点小权力的快感中。</UserCard>
    </div>
    <div class="col-sm-6 col-xl-4">
      <UserCard>耳朵听力分辨率较差，在家办公和同事电话会议基本听不清在说什么，痛苦啊！</UserCard>
    </div>
  </div>
</template>
<script>
import UserCard from "./comps/UserCard";
export default {
  components: {
    UserCard
  }
};
</script>
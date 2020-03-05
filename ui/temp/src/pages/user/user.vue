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
    <div class="col-sm-6 col-xl-4" v-for="item of admins" :key="item.id">
      <UserCard :admin="item"></UserCard>
    </div>
  </div>
</template>
<script>
import {getAllAdmin} from "@/api/admin"
import UserCard from "./comps/UserCard";
export default {
  components: {
    UserCard
  },
  data() {
    return {
      admins: []
    }
  },
  mounted() {
    this.loadData()
  },
  methods: {
    loadData() {
      getAllAdmin().then(res => {
        this.admins = res.admins
      })
    }
  }
};
</script>
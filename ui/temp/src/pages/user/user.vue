<template>
  <div class="row">
    <div class="col-12">
      <TBlock>
        <div class="row">
          <div class="col-6 push">
            <b-input-group size="sm">
              <b-form-input></b-form-input>
              <b-input-group-append>
                <b-button variant="outline-secondary"><TIcon icon="search" /></b-button>
              </b-input-group-append>
            </b-input-group>
          </div>
          <div class="col-6 push">
            <b-button variant="success" size="sm" @click="userlVisit=true"><TIcon icon="plus" /> 创建新用户</b-button>
            <b-button variant="secondary" class="ml-1" size="sm" @click="$router.push('/role')"><TIcon icon="user" pack="si" /> 角色管理</b-button>
          </div>
        </div>
      </TBlock>
    </div>
    <div class="col-sm-6 col-xl-4" v-for="item of admins" :key="item.id">
      <UserCard :admin="item" @onedit="oneditHandler"></UserCard>
    </div>
    <b-modal
      v-model="userlVisit"
      size="lg"
      title="设置角色"
      no-close-on-backdrop
      @ok="submitUser"
      @hide="userModalHide"
      ok-title="保存"
      cancel-title="关闭"
    >
      <UserModal ref="userMoadl" :editId="editId" @savedUser="savedUserHandler"></UserModal>
    </b-modal>
  </div>
</template>
<script>
import {getAllAdmin} from "@/api/admin"
import UserCard from "./comps/UserCard";
import UserModal from "./comps/UserModal";

export default {
  components: {
    UserCard,
    UserModal
  },
  data() {
    return {
      editId: 0,
      userlVisit: false,
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
    },
    submitUser(bvModalEvt) {
      bvModalEvt.preventDefault()
      this.$refs.userMoadl.submit()
    },
    savedUserHandler() {
      this.loadData()
      this.userlVisit = false
    },
    oneditHandler(id) {
      this.userlVisit = true
      this.editId = id
    },
    userModalHide() {
      this.editId = 0
    }
  }
};
</script>
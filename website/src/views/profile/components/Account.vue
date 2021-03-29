<template>
  <el-form>
    <el-form-item label="Name">
      <el-input v-model.trim="user.name" />
    </el-form-item>
    <el-form-item label="Email">
      <el-input v-model.trim="user.email" />
    </el-form-item>
    <el-form-item label="Password">
      <el-input v-model.trim="user.password" show-password />
    </el-form-item>
    <el-form-item>
      <el-button type="primary" @click="submit">Update</el-button>
    </el-form-item>
  </el-form>
</template>

<script>
import { updateUser } from '@/api/user'
export default {
  props: {
    user: {
      type: Object,
      default: () => {
        return {
          name: '',
          email: '',
          password: ''
        }
      }
    }
  },
  methods: {
    submit() {
      updateUser(this.user).then(response => {
        console.log(this.$store)
        this.$store.commit('user/SET_NAME', this.user.name)
        this.$store.commit('user/SET_AVATAR', this.user.avatar)
        this.$store.commit('user/SET_EMAIL', this.user.email)
        this.$message({
          message: 'User information has been updated successfully',
          type: 'success',
          duration: 5 * 1000
        })
      })
    }
  }
}
</script>

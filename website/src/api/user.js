import request from '@/utils/request'

export function login(data) {
  return request({
    url: '/login',
    method: 'post',
    data
  })
}

export function getInfo() {
  return request({
    url: '/user_info',
    method: 'get'
  })
}

export function updateUser(data) {
  return request({
    url: '/user_update',
    method: 'put',
    data
  })
}

export function logout() {
  return request({
    url: '/logout',
    method: 'post'
  })
}

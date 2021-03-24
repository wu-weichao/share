import request from '@/utils/request'

export function fetchList(query) {
  return request({
    url: '/topics',
    method: 'get',
    params: query
  })
}

export function fetchTopic(id) {
  return request({
    url: '/topics/' + id,
    method: 'get'
  })
}

export function createTopic(data) {
  return request({
    url: '/topics',
    method: 'post',
    data
  })
}

export function updateTopic(id, data) {
  return request({
    url: '/topics/' + id,
    method: 'put',
    data
  })
}

export function deleteTopic(id) {
  return request({
    url: '/topics/' + id,
    method: 'delete'
  })
}

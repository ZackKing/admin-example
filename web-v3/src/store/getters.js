const getters = {
  sidebar: state => state.app.sidebar,
  device: state => state.app.device,
  token: state => state.user.token,
  avatar: state => state.user.avatar,
  uid: state => state.user.uid,
  name: state => state.user.name,
  roles: state => state.user.roles,
  permission_routes: state => state.permission.routes,
  accessUrls: state => state.permission.accessUrls,
  platform: state => state.common.platform,
  accessRoutes: state => state.common.accessRoutes
}
export default getters

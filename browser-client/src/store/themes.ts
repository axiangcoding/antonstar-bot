const anton_star = {
	common: {
		headerHeight: '46px',
		footerHeight: '100px',
	},
	Layout: {
		headerTextColor: '#d4d4d4',
		headerTransparentColor: '#40485bdd',
		headerNavHoverColor: '#f2f2f2',
		headerNavActiveColor: '#fff',
		headerColor: '#40485b',
	},
	Button: {
		headerHoverColor: '#eee',
		headerAcitveColor: '#fff',
	},
}

export default {
	namespaced: true,
	state: {
		themesOverides: {
			anton_star,
		},
  },
}

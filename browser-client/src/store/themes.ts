const anton_star = {
	common: {
		headerHeight: '52px',
		footerHeight: '100px',
		primaryColor: '#FCA706',
		primaryColorHover: '#F26E1D',
		primaryColorPressed: '#FC5C04',
	},
	Layout: {
		headerTextColor: '#d4d4d4',
		headerTransparentColor: '#40485bdd',
		primaryTransparent: '#FCA706dd',
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

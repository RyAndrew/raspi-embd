/*
 * File: app/view/ServoTesterPanel.js
 *
 * This file was generated by Sencha Architect
 * http://www.sencha.com/products/architect/
 *
 * This file requires use of the Ext JS 6.5.x Modern library, under independent license.
 * License of Sencha Architect does not include license for Ext JS 6.5.x Modern. For more
 * details see http://www.sencha.com/license or contact license@sencha.com.
 *
 * This file will be auto-generated each and everytime you save your project.
 *
 * Do NOT hand edit this file.
 */

Ext.define('ServoTester.view.ServoTesterPanel', {
	extend: 'Ext.Panel',
	alias: 'widget.servotesterpanel',

	requires: [
		'ServoTester.view.ServoTesterPanelViewModel',
		'Ext.Container',
		'Ext.field.Slider',
		'Ext.field.Number'
	],

	viewModel: {
		type: 'servotesterpanel'
	},
	fullscreen: true,
	padding: 10,
	layout: 'vbox',
	title: 'Servo Tester',
	defaultListenerScope: true,

	items: [
		{
			xtype: 'container',
			height: 25,
			itemId: 'headerContainer',
			layout: 'hbox',
			items: [
				{
					xtype: 'container',
					flex: 1,
					html: 'PWM Channel #'
				},
				{
					xtype: 'container',
					width: 80,
					html: 'Value'
				}
			]
		},
		{
			xtype: 'container',
			itemId: 'servo0',
			layout: 'hbox',
			items: [
				{
					xtype: 'sliderfield',
					flex: 1,
					bind: '{servo0value}',
					name: 'servo0',
					label: '0',
					labelWidth: 30,
					listeners: {
						drag: 'onMysliderfieldDrag0',
						change: 'onMysliderfieldChange0'
					}
				},
				{
					xtype: 'numberfield',
					bind: '{servo0value}',
					width: 80,
					margin: '0 10 0 10',
					clearable: false,
					textAlign: 'center',
					maxValue: 100,
					minValue: 0
				}
			]
		},
		{
			xtype: 'container',
			itemId: 'servo1',
			layout: 'hbox',
			items: [
				{
					xtype: 'sliderfield',
					flex: 1,
					bind: '{servo1value}',
					name: 'servo1',
					label: '1',
					labelWidth: 30,
					listeners: {
						drag: 'onMysliderfieldDrag1',
						change: 'onMysliderfieldChange1'
					}
				},
				{
					xtype: 'numberfield',
					bind: '{servo1value}',
					width: 80,
					margin: '0 10 0 10',
					clearable: false,
					textAlign: 'center',
					maxValue: 100,
					minValue: 0
				}
			]
		},
		{
			xtype: 'container',
			itemId: 'servo2',
			layout: 'hbox',
			items: [
				{
					xtype: 'sliderfield',
					flex: 1,
					bind: '{servo2value}',
					name: 'servo2',
					label: '2',
					labelWidth: 30,
					listeners: {
						drag: 'onMysliderfieldDrag2',
						change: 'onMysliderfieldChange2'
					}
				},
				{
					xtype: 'numberfield',
					bind: '{servo2value}',
					width: 80,
					margin: '0 10 0 10',
					clearable: false,
					textAlign: 'center',
					maxValue: 100,
					minValue: 0
				}
			]
		},
		{
			xtype: 'container',
			itemId: 'servo3',
			layout: 'hbox',
			items: [
				{
					xtype: 'sliderfield',
					flex: 1,
					bind: '{servo3value}',
					name: 'servo3',
					label: '3',
					labelWidth: 30,
					listeners: {
						drag: 'onMysliderfieldDrag3',
						change: 'onMysliderfieldChange3'
					}
				},
				{
					xtype: 'numberfield',
					bind: '{servo3value}',
					width: 80,
					margin: '0 10 0 10',
					clearable: false,
					textAlign: 'center',
					maxValue: 100,
					minValue: 0
				}
			]
		},
		{
			xtype: 'container',
			itemId: 'servo4',
			layout: 'hbox',
			items: [
				{
					xtype: 'sliderfield',
					flex: 1,
					bind: '{servo4value}',
					name: 'servo4',
					label: '4',
					labelWidth: 30,
					listeners: {
						drag: 'onMysliderfieldDrag4',
						change: 'onMysliderfieldChange4'
					}
				},
				{
					xtype: 'numberfield',
					bind: '{servo4value}',
					width: 80,
					margin: '0 10 0 10',
					clearable: false,
					textAlign: 'center',
					maxValue: 100,
					minValue: 0
				}
			]
		},
		{
			xtype: 'container',
			itemId: 'servo5',
			layout: 'hbox',
			items: [
				{
					xtype: 'sliderfield',
					flex: 1,
					bind: '{servo5value}',
					name: 'servo5',
					label: '5',
					labelWidth: 30,
					listeners: {
						drag: 'onMysliderfieldDrag5',
						change: 'onMysliderfieldChange5'
					}
				},
				{
					xtype: 'numberfield',
					bind: '{servo5value}',
					width: 80,
					margin: '0 10 0 10',
					clearable: false,
					textAlign: 'center',
					maxValue: 100,
					minValue: 0
				}
			]
		},
		{
			xtype: 'container',
			itemId: 'servo6',
			layout: 'hbox',
			items: [
				{
					xtype: 'sliderfield',
					flex: 1,
					bind: '{servo6value}',
					name: 'servo6',
					label: '6',
					labelWidth: 30,
					listeners: {
						drag: 'onMysliderfieldDrag6',
						change: 'onMysliderfieldChange6'
					}
				},
				{
					xtype: 'numberfield',
					bind: '{servo6value}',
					width: 80,
					margin: '0 10 0 10',
					clearable: false,
					textAlign: 'center',
					maxValue: 100,
					minValue: 0
				}
			]
		},
		{
			xtype: 'container',
			itemId: 'servo7',
			layout: 'hbox',
			items: [
				{
					xtype: 'sliderfield',
					flex: 1,
					bind: '{servo7value}',
					name: 'servo7',
					label: '7',
					labelWidth: 30,
					listeners: {
						drag: 'onMysliderfieldDrag7',
						change: 'onMysliderfieldChange7'
					}
				},
				{
					xtype: 'numberfield',
					bind: '{servo7value}',
					width: 80,
					margin: '0 10 0 10',
					clearable: false,
					textAlign: 'center',
					maxValue: 100,
					minValue: 0
				}
			]
		},
		{
			xtype: 'container',
			itemId: 'servo8',
			layout: 'hbox',
			items: [
				{
					xtype: 'sliderfield',
					flex: 1,
					bind: '{servo8value}',
					name: 'servo8',
					label: '8',
					labelWidth: 30,
					listeners: {
						drag: 'onMysliderfieldDrag8',
						change: 'onMysliderfieldChange8'
					}
				},
				{
					xtype: 'numberfield',
					bind: '{servo8value}',
					width: 80,
					margin: '0 10 0 10',
					clearable: false,
					textAlign: 'center',
					maxValue: 100,
					minValue: 0
				}
			]
		},
		{
			xtype: 'container',
			itemId: 'servo9',
			layout: 'hbox',
			items: [
				{
					xtype: 'sliderfield',
					flex: 1,
					bind: '{servo9value}',
					name: 'servo9',
					label: '9',
					labelWidth: 30,
					listeners: {
						drag: 'onMysliderfieldDrag9',
						change: 'onMysliderfieldChange9'
					}
				},
				{
					xtype: 'numberfield',
					bind: '{servo9value}',
					width: 80,
					margin: '0 10 0 10',
					clearable: false,
					textAlign: 'center',
					maxValue: 100,
					minValue: 0
				}
			]
		},
		{
			xtype: 'container',
			itemId: 'servo10',
			layout: 'hbox',
			items: [
				{
					xtype: 'sliderfield',
					flex: 1,
					bind: '{servo10value}',
					name: 'servo10',
					label: '10',
					labelWidth: 30,
					listeners: {
						drag: 'onMysliderfieldDrag10',
						change: 'onMysliderfieldChange10'
					}
				},
				{
					xtype: 'numberfield',
					bind: '{servo10value}',
					width: 80,
					margin: '0 10 0 10',
					clearable: false,
					textAlign: 'center',
					maxValue: 100,
					minValue: 0
				}
			]
		},
		{
			xtype: 'container',
			itemId: 'servo11',
			layout: 'hbox',
			items: [
				{
					xtype: 'sliderfield',
					flex: 1,
					bind: '{servo11value}',
					name: 'servo11',
					label: '11',
					labelWidth: 30,
					listeners: {
						drag: 'onMysliderfieldDrag11',
						change: 'onMysliderfieldChange11'
					}
				},
				{
					xtype: 'numberfield',
					bind: '{servo11value}',
					width: 80,
					margin: '0 10 0 10',
					clearable: false,
					textAlign: 'center',
					maxValue: 100,
					minValue: 0
				}
			]
		},
		{
			xtype: 'container',
			itemId: 'servo12',
			layout: 'hbox',
			items: [
				{
					xtype: 'sliderfield',
					flex: 1,
					bind: '{servo12value}',
					name: 'servo12',
					label: '12',
					labelWidth: 30,
					listeners: {
						drag: 'onMysliderfieldDrag12',
						change: 'onMysliderfieldChange12'
					}
				},
				{
					xtype: 'numberfield',
					bind: '{servo12value}',
					width: 80,
					margin: '0 10 0 10',
					clearable: false,
					textAlign: 'center',
					maxValue: 100,
					minValue: 0
				}
			]
		},
		{
			xtype: 'container',
			itemId: 'servo13',
			layout: 'hbox',
			items: [
				{
					xtype: 'sliderfield',
					flex: 1,
					bind: '{servo13value}',
					name: 'servo13',
					label: '13',
					labelWidth: 30,
					listeners: {
						drag: 'onMysliderfieldDrag13',
						change: 'onMysliderfieldChange13'
					}
				},
				{
					xtype: 'numberfield',
					bind: '{servo13value}',
					width: 80,
					margin: '0 10 0 10',
					clearable: false,
					textAlign: 'center',
					maxValue: 100,
					minValue: 0
				}
			]
		},
		{
			xtype: 'container',
			itemId: 'servo14',
			layout: 'hbox',
			items: [
				{
					xtype: 'sliderfield',
					flex: 1,
					bind: '{servo14value}',
					name: 'servo14',
					label: '14',
					labelWidth: 30,
					listeners: {
						drag: 'onMysliderfieldDrag14',
						change: 'onMysliderfieldChange14'
					}
				},
				{
					xtype: 'numberfield',
					bind: '{servo14value}',
					width: 80,
					margin: '0 10 0 10',
					clearable: false,
					textAlign: 'center',
					maxValue: 100,
					minValue: 0
				}
			]
		},
		{
			xtype: 'container',
			itemId: 'servo15',
			layout: 'hbox',
			items: [
				{
					xtype: 'sliderfield',
					flex: 1,
					bind: '{servo15value}',
					name: 'servo15',
					label: '15',
					labelWidth: 30,
					listeners: {
						drag: 'onMysliderfieldDrag15',
						change: 'onMysliderfieldChange15'
					}
				},
				{
					xtype: 'numberfield',
					bind: '{servo15value}',
					width: 80,
					margin: '0 10 0 10',
					clearable: false,
					textAlign: 'center',
					maxValue: 100,
					minValue: 0
				}
			]
		}
	],

	onMysliderfieldDrag0: function(sliderfield, sl, thumb, e, eOpts) {
		this.fieldChange(sliderfield.name,e[0]);

	},

	onMysliderfieldChange0: function(me, newValue, oldValue, eOpts) {
		 this.fieldChange(me.name, newValue);
	},

	onMysliderfieldDrag1: function(sliderfield, sl, thumb, e, eOpts) {
		this.fieldChange(sliderfield.name,e[0]);

	},

	onMysliderfieldChange1: function(me, newValue, oldValue, eOpts) {
		 this.fieldChange(me.name, newValue);
	},

	onMysliderfieldDrag2: function(sliderfield, sl, thumb, e, eOpts) {
		this.fieldChange(sliderfield.name,e[0]);

	},

	onMysliderfieldChange2: function(me, newValue, oldValue, eOpts) {
		 this.fieldChange(me.name, newValue);
	},

	onMysliderfieldDrag3: function(sliderfield, sl, thumb, e, eOpts) {
		this.fieldChange(sliderfield.name,e[0]);

	},

	onMysliderfieldChange3: function(me, newValue, oldValue, eOpts) {
		 this.fieldChange(me.name, newValue);
	},

	onMysliderfieldDrag4: function(sliderfield, sl, thumb, e, eOpts) {
		this.fieldChange(sliderfield.name,e[0]);

	},

	onMysliderfieldChange4: function(me, newValue, oldValue, eOpts) {
		 this.fieldChange(me.name, newValue);
	},

	onMysliderfieldDrag5: function(sliderfield, sl, thumb, e, eOpts) {
		this.fieldChange(sliderfield.name,e[0]);

	},

	onMysliderfieldChange5: function(me, newValue, oldValue, eOpts) {
		 this.fieldChange(me.name, newValue);
	},

	onMysliderfieldDrag6: function(sliderfield, sl, thumb, e, eOpts) {
		this.fieldChange(sliderfield.name,e[0]);

	},

	onMysliderfieldChange6: function(me, newValue, oldValue, eOpts) {
		 this.fieldChange(me.name, newValue);
	},

	onMysliderfieldDrag7: function(sliderfield, sl, thumb, e, eOpts) {
		this.fieldChange(sliderfield.name,e[0]);

	},

	onMysliderfieldChange7: function(me, newValue, oldValue, eOpts) {
		 this.fieldChange(me.name, newValue);
	},

	onMysliderfieldDrag8: function(sliderfield, sl, thumb, e, eOpts) {
		this.fieldChange(sliderfield.name,e[0]);

	},

	onMysliderfieldChange8: function(me, newValue, oldValue, eOpts) {
		 this.fieldChange(me.name, newValue);
	},

	onMysliderfieldDrag9: function(sliderfield, sl, thumb, e, eOpts) {
		this.fieldChange(sliderfield.name,e[0]);

	},

	onMysliderfieldChange9: function(me, newValue, oldValue, eOpts) {
		 this.fieldChange(me.name, newValue);
	},

	onMysliderfieldDrag10: function(sliderfield, sl, thumb, e, eOpts) {
		this.fieldChange(sliderfield.name,e[0]);

	},

	onMysliderfieldChange10: function(me, newValue, oldValue, eOpts) {
		 this.fieldChange(me.name, newValue);
	},

	onMysliderfieldDrag11: function(sliderfield, sl, thumb, e, eOpts) {
		this.fieldChange(sliderfield.name,e[0]);

	},

	onMysliderfieldChange11: function(me, newValue, oldValue, eOpts) {
		 this.fieldChange(me.name, newValue);
	},

	onMysliderfieldDrag12: function(sliderfield, sl, thumb, e, eOpts) {
		this.fieldChange(sliderfield.name,e[0]);

	},

	onMysliderfieldChange12: function(me, newValue, oldValue, eOpts) {
		 this.fieldChange(me.name, newValue);
	},

	onMysliderfieldDrag13: function(sliderfield, sl, thumb, e, eOpts) {
		this.fieldChange(sliderfield.name,e[0]);

	},

	onMysliderfieldChange13: function(me, newValue, oldValue, eOpts) {
		 this.fieldChange(me.name, newValue);
	},

	onMysliderfieldDrag14: function(sliderfield, sl, thumb, e, eOpts) {
		this.fieldChange(sliderfield.name,e[0]);

	},

	onMysliderfieldChange14: function(me, newValue, oldValue, eOpts) {
		 this.fieldChange(me.name, newValue);
	},

	onMysliderfieldDrag15: function(sliderfield, sl, thumb, e, eOpts) {
		this.fieldChange(sliderfield.name,e[0]);

	},

	onMysliderfieldChange15: function(me, newValue, oldValue, eOpts) {
		 this.fieldChange(me.name, newValue);
	},

	fieldChange: function(name, value) {
		console.log(name+' = '+value);

		if(this.activeRequest){
		    Ext.Ajax.abort(this.activeRequest);
		}

		this.activeRequest = AERP.Ajax.request({
		    url: '/updateServo',
		    params:{
		        servo: name,
		        value: value
		    },
		    success:function(response){
		        this.activeRequest = false;
		    },
		    failure:function(){

		    },
		    scope:this
		});

	}

});
/*
 * File: app/view/RcSteeringPanel.js
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

Ext.define('RcSteering.view.RcSteeringPanel', {
    extend: 'Ext.form.Panel',
    alias: 'widget.rcsteeringpanel',

    requires: [
        'RcSteering.view.RcSteeringPanelViewModel',
        'Ext.Container',
        'Ext.field.Slider',
        'Ext.field.Number',
        'Ext.Button'
    ],

    viewModel: {
        type: 'rcsteeringpanel'
    },
    fullscreen: true,
    title: 'RC Steering Tester',
    defaultListenerScope: true,

    items: [
        {
            xtype: 'container',
            userCls: 'steering-label',
            margin: 5,
            layout: 'hbox',
            items: [
                {
                    xtype: 'container',
                    width: 30,
                    html: 'L',
                    margin: '0 0 0 60'
                },
                {
                    xtype: 'container',
                    flex: 1
                },
                {
                    xtype: 'container',
                    width: 30,
                    html: 'R'
                }
            ]
        },
        {
            xtype: 'sliderfield',
            bind: '{steeringSetValue}',
            itemId: 'steeringSetPointSlider',
            name: 'steeringSlider',
            margin: '0 10 5 10',
            label: 'Steer',
            labelWidth: 50,
            value: 50,
            liveUpdate: true
        },
        {
            xtype: 'container',
            itemId: 'hbox',
            layout: 'hbox',
            items: [
                {
                    xtype: 'container',
                    itemId: 'setPointHbox',
                    userCls: 'variable-box',
                    items: [
                        {
                            xtype: 'container',
                            itemId: 'setPointActualLabel',
                            html: 'Set Point & Actual',
                            margin: 10
                        },
                        {
                            xtype: 'numberfield',
                            bind: '{steeringSetValue}',
                            itemId: 'steeringSetPointText',
                            width: 150,
                            margin: 10,
                            label: 'set point',
                            labelWidth: 70,
                            value: 50,
                            clearable: false,
                            maxValue: 100,
                            minValue: 0,
                            listeners: {
                                change: 'onSteeringSetPointTextChange'
                            }
                        },
                        {
                            xtype: 'textfield',
                            itemId: 'steeringCurrent',
                            width: 150,
                            margin: 10,
                            label: 'current',
                            labelWidth: 70,
                            value: 0,
                            clearable: false
                        },
                        {
                            xtype: 'textfield',
                            itemId: 'pidError',
                            width: 150,
                            margin: 10,
                            label: 'error',
                            labelWidth: 70,
                            value: 0,
                            clearable: false
                        }
                    ]
                },
                {
                    xtype: 'container',
                    itemId: 'pidHbox',
                    userCls: 'variable-box',
                    items: [
                        {
                            xtype: 'container',
                            itemId: 'pidVariablesLabel',
                            html: 'PID Variables',
                            margin: 10
                        },
                        {
                            xtype: 'textfield',
                            itemId: 'pidConstantP',
                            width: 100,
                            margin: 10,
                            label: 'P',
                            labelWidth: 25,
                            value: 0,
                            clearable: false
                        },
                        {
                            xtype: 'textfield',
                            itemId: 'pidConstantI',
                            width: 100,
                            margin: 10,
                            label: 'I',
                            labelWidth: 25,
                            value: 0,
                            clearable: false
                        },
                        {
                            xtype: 'textfield',
                            itemId: 'pidConstantD',
                            width: 100,
                            margin: 10,
                            label: 'D',
                            labelWidth: 25,
                            value: 0,
                            clearable: false
                        }
                    ]
                },
                {
                    xtype: 'container',
                    itemId: 'pidOutputHbox',
                    userCls: 'variable-box',
                    items: [
                        {
                            xtype: 'container',
                            itemId: 'pidOutputsLabel',
                            html: 'PID Outputs',
                            margin: 10
                        },
                        {
                            xtype: 'textfield',
                            itemId: 'pidIntegral',
                            width: 150,
                            margin: 10,
                            label: 'integral',
                            labelWidth: 60,
                            clearable: false
                        },
                        {
                            xtype: 'textfield',
                            itemId: 'pidOutput',
                            width: 150,
                            margin: 10,
                            label: 'output',
                            labelWidth: 60,
                            clearable: false
                        }
                    ]
                },
                {
                    xtype: 'container',
                    itemId: 'buttonsHbox',
                    userCls: 'variable-box',
                    layout: 'vbox',
                    items: [
                        {
                            xtype: 'container',
                            itemId: 'buttonsLabel',
                            html: 'Control',
                            margin: 10
                        },
                        {
                            xtype: 'button',
                            margin: '10 20 10 20',
                            text: 'Stop',
                            listeners: {
                                tap: 'onMybuttonTap'
                            }
                        },
                        {
                            xtype: 'button',
                            margin: '10 20 10 20',
                            text: 'Stop Steering Loop',
                            listeners: {
                                tap: 'onMybuttonTap1'
                            }
                        },
                        {
                            xtype: 'button',
                            margin: '10 20 10 20',
                            text: 'Start Steering Loop',
                            listeners: {
                                tap: 'onMybuttonTap11'
                            }
                        }
                    ]
                }
            ]
        },
        {
            xtype: 'container',
            itemId: 'debugOutputLabel',
            html: 'Debug Output',
            margin: 10
        },
        {
            xtype: 'container',
            height: 240,
            itemId: 'debugOutputContainerOuter',
            userCls: 'debugOutputContainer',
            margin: 10,
            scrollable: 'vertical',
            layout: 'vbox',
            items: [
                {
                    xtype: 'container',
                    itemId: 'debugOutputContainer'
                }
            ]
        }
    ],
    listeners: {
        painted: 'onFormpanelPainted'
    },

    onSteeringSetPointTextChange: function(field, newValue, oldValue, eOpts) {
        newValue = newValue[0];

        if(!this.lastSteeringChangeDefered){
            console.log("Steering Set = "+newValue);
            this.lastSteeringChangeDefered = true;

            if(this.lastSent != newValue){
                this.sendUpdateSteering(newValue);
                this.lastSent = newValue;
            }
            Ext.defer(function(){
                this.lastSteeringChangeDefered = false;
                var currentValue = field.getValue()[0];
                if(this.lastSent != currentValue){
                    this.lastSent = currentValue;
                    this.sendUpdateSteering(currentValue);
                }
            },100,this);
        }
    },

    onMybuttonTap: function(button, e, eOpts) {
        this.stopSteeringMovement();
    },

    onMybuttonTap1: function(button, e, eOpts) {
        this.socketSend(Ext.encode({
            action:'stopSteeringControlLoop'
        }));
    },

    onMybuttonTap11: function(button, e, eOpts) {
        this.socketSend(Ext.encode({
            action:'startSteeringControlLoop'
        }));
    },

    onFormpanelPainted: function(sender, element, eOpts) {
        this.socketInit();
    },

    socketInit: function() {
        	this.socket = new WebSocket('ws://'+window.location.host+'/wsapi');

            this.appendDebugOutput("Opening Websocket");

        	this.socket.onerror = (function(){
        		this.appendDebugOutput("error with websocket!");
                this.socket = false;
        		return;
        	}).bind(this);
        	this.socket.onclose = (function(){
        		this.appendDebugOutput("websocket closed!");
                this.socket = false;
        		return;
        	}).bind(this);
        	this.socket.onmessage = this.socketReceive.bind(this);
    },

    socketSend: function(message) {
        if(!this.socket){
            console.log('Websocket not started or failed');
            return false;
        }

        if(this.socket.readyState !== WebSocket.OPEN){
            console.log('Websocket not connected');
            return false;
        }

        this.socket.send(message);
    },

    stopSteeringMovement: function() {
        this.socketSend(Ext.encode({
            action:'stopSteeringMovement'
        }));
    },

    socketReceive: function(message) {
        // console.log('recieved message');
        // console.log(message.data);

        this.appendDebugOutput('recieved message: '+message.data);

        var jsonData = JSON.parse(message.data);

        this.queryById('steeringCurrent').setValue(jsonData.steeringCurrent);
    },

    sendUpdateSteering: function(value) {
        this.socketSend(Ext.encode({
            action:'updateSteering',
            value:value
        }));
    },

    appendDebugOutput: function(message) {
        // console.log(message);
        // console.log(this.queryById('debugOutputContainer'));
        var dom = this.queryById('debugOutputContainer').el.dom;
        // console.log(dom);
        dom.innerHTML += message + "<BR>\r\n";
        // console.log(this.queryById('debugOutputContainerOuter'));
        containerDom = this.queryById('debugOutputContainerOuter').bodyElement.dom;
        containerDom.scrollTop = dom.clientHeight;
    }

});
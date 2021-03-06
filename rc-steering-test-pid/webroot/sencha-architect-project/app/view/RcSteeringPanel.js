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
        'Ext.field.Slider',
        'Ext.field.Number',
        'Ext.Button',
        'Ext.chart.CartesianChart',
        'Ext.chart.axis.Numeric',
        'Ext.chart.series.Line'
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
                    html: 'Left',
                    margin: '0 0 0 60'
                },
                {
                    xtype: 'container',
                    flex: 1
                },
                {
                    xtype: 'container',
                    html: 'Right'
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
            labelWidth: 65,
            value: 50,
            liveUpdate: true
        },
        {
            xtype: 'container',
            userCls: 'steering-label',
            margin: 5,
            layout: 'hbox',
            items: [
                {
                    xtype: 'container',
                    html: 'Reverse',
                    margin: '0 0 0 60'
                },
                {
                    xtype: 'container',
                    flex: 1
                },
                {
                    xtype: 'container',
                    html: 'Forward'
                }
            ]
        },
        {
            xtype: 'sliderfield',
            bind: '{throttleSetValue}',
            itemId: 'throttleSlider',
            name: 'throttleSlider',
            margin: '0 10 5 10',
            label: 'Throttle',
            labelWidth: 65,
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
                    itemId: 'setPointHbox1',
                    userCls: 'variable-box',
                    items: [
                        {
                            xtype: 'container',
                            itemId: 'throttleLabel',
                            html: 'Throttle',
                            margin: 10
                        },
                        {
                            xtype: 'numberfield',
                            bind: '{throttleSetValue}',
                            itemId: 'throttleText',
                            width: 150,
                            margin: 10,
                            label: 'throttle',
                            labelWidth: 70,
                            value: 0,
                            clearable: false,
                            maxValue: 100,
                            minValue: 0,
                            listeners: {
                                change: 'onThrottleTextChange'
                            }
                        }
                    ]
                },
                {
                    xtype: 'container',
                    itemId: 'setPointHbox',
                    userCls: 'variable-box',
                    items: [
                        {
                            xtype: 'container',
                            itemId: 'setPointActualLabel',
                            html: 'Steer Set & Actual',
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
                            html: 'PID Constants',
                            margin: 10
                        },
                        {
                            xtype: 'numberfield',
                            itemId: 'pidConstantP',
                            width: 100,
                            margin: 10,
                            label: 'P',
                            labelWidth: 25,
                            value: 0,
                            clearable: false
                        },
                        {
                            xtype: 'numberfield',
                            itemId: 'pidConstantI',
                            width: 100,
                            margin: 10,
                            label: 'I',
                            labelWidth: 25,
                            value: 0,
                            clearable: false
                        },
                        {
                            xtype: 'numberfield',
                            itemId: 'pidConstantD',
                            width: 100,
                            margin: '10 10 0 10',
                            label: 'D',
                            labelWidth: 25,
                            value: 0,
                            clearable: false
                        },
                        {
                            xtype: 'button',
                            itemId: 'updatePidConstants',
                            margin: '10 10 10 30',
                            text: 'Update',
                            listeners: {
                                tap: 'onMybutton1Tap'
                            }
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
                            margin: '0 20 5 20',
                            text: 'Stop Steering',
                            listeners: {
                                tap: 'onMybuttonTap'
                            }
                        },
                        {
                            xtype: 'button',
                            margin: '5 20 5 20',
                            text: 'Stop Throttle',
                            listeners: {
                                tap: 'onMybuttonTap2'
                            }
                        },
                        {
                            xtype: 'button',
                            margin: '5 20 5 20',
                            text: 'Stop Steering Loop',
                            listeners: {
                                tap: 'onMybuttonTap1'
                            }
                        },
                        {
                            xtype: 'button',
                            margin: '5 20 5 20',
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
            xtype: 'cartesian',
            height: 250,
            itemId: 'steeringChart',
            animation: false,
            colors: [
                '#115fa6',
                '#94ae0a',
                '#a61120',
                '#ff8809',
                '#ffd13e',
                '#a61187',
                '#24ad9a',
                '#7c7474',
                '#a66111'
            ],
            bind: {
                store: '{steeringChartStore}'
            },
            axes: [
                {
                    type: 'numeric',
                    grid: {
                        odd: {
                            fill: '#e8e8e8'
                        }
                    },
                    maximum: 1800,
                    minimum: 0,
                    position: 'left',
                    title: 'Position'
                }
            ],
            series: [
                {
                    type: 'line',
                    colors: [
                        'rgba(200,0,0,0.3)'
                    ],
                    style: {
                        stroke: 'rgb(200,0,0)',
                        step: true
                    },
                    xField: 'x',
                    yField: 'steeringCurrent'
                },
                {
                    type: 'line',
                    colors: [
                        'rgba(0,200,0,0.3)'
                    ],
                    style: {
                        stroke: 'rgb(0,200,0)',
                        step: true
                    },
                    xField: 'x',
                    yField: 'steeringTargetPoint'
                }
            ]
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

    onThrottleTextChange: function(field, newValue, oldValue, eOpts) {
        console.log(newValue);
        if(newValue && newValue.constructor === Array){
            newValue = newValue[0];
        }

        this.sendUpdateThrottle(newValue);
    },

    onSteeringSetPointTextChange: function(field, newValue, oldValue, eOpts) {
        console.log(newValue);
        if(newValue && newValue.constructor === Array){
            newValue = newValue[0];
        }

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
            },2,this);
        }
    },

    onMybutton1Tap: function(button, e, eOpts) {
        var p = this.queryById('pidConstantP').getValue(),
            i = this.queryById('pidConstantI').getValue(),
            d = this.queryById('pidConstantD').getValue();

        if(p === null || i === null || d === null){
            Ext.Msg.alert(' ','Invalid Constants! Please try again!');
            return false;
        }
        this.socketSend(Ext.encode({
            action:'updatePidConstants',
            p: p,
            i: i,
            d: d
        }));
    },

    onMybuttonTap: function(button, e, eOpts) {
        this.stopSteeringMovement();
    },

    onMybuttonTap2: function(button, e, eOpts) {
        this.socketSend(Ext.encode({
            action:'stopThrottle'
        }));
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

            this.recievedMessages = 0;
            this.chartedMessages = 0;

            this.chartArray = [];
            this.chartStore = this.lookupViewModel().getStore('steeringChartStore');
            this.steeringChart = this.queryById('steeringChart');

            this.appendDebugOutput("Opening Websocket");

        	this.socket.onerror = (function(){
        		this.appendDebugOutput("error with websocket!");
                this.socket = false;
        	}).bind(this);
        	this.socket.onclose = (function(){
        		this.appendDebugOutput("websocket closed!");
                this.socket = false;
        	}).bind(this);

        	this.socket.onmessage = this.socketReceive.bind(this);


        	this.socket.onopen = (function(){
                this.socketSend(Ext.encode({
                    action:'getPidConstants'
                }));
        	}).bind(this);
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
        //this.appendDebugOutput('recieved message: '+message.data);


        var jsonData = JSON.parse(message.data);

        if(!jsonData.msgType){
            console.log('Missing msgType!');
            console.log(jsonData);
            return false;
        }
        switch(jsonData.msgType){
            default:
                console.log('Invalid msgType!');
                console.log(jsonData);
                break;
            case 'status':
                this.msgUpdateStatus(jsonData);
                break;
            case 'pidConstants':
                this.msgUpdatePidConstants(jsonData);
                break;
        }


    },

    sendUpdateSteering: function(value) {
        this.socketSend(Ext.encode({
            action:'updateSteering',
            value:value
        }));
    },

    sendUpdateThrottle: function(value) {
        this.socketSend(Ext.encode({
            action:'updateThrottle',
            value:value
        }));
    },

    msgUpdateStatus: function(jsonData) {
        // this.recievedMessages++;
        this.chartedMessages++;


        // if(this.recievedMessages > 100){
        //     this.recievedMessages = 0;
        //     this.clearDebugOutput();
        //     //chartStore.removeAll();
        // }



        if(this.chartStore.count() > 100){
            this.chartStore.removeAt(0,1);
        }

        //this.steeringChart.suspendAnimation();
        // if(this.chartArray.length > 100){
        //     this.chartArray.shift();
        // }
        // this.chartArray.push([this.chartedMessages, jsonData.steeringCurrent, jsonData.steeringTargetPoint]);

        // this.chartStore.loadData(this.chartArray);

        this.chartStore.loadData([[this.chartedMessages, jsonData.steeringCurrent, jsonData.steeringTargetPoint]], true);

        //this.steeringChart.resumeAnimation();

        this.queryById('steeringCurrent').setValue(jsonData.steeringCurrent);
        this.queryById('pidError').setValue(jsonData.pidError);
    },

    msgUpdatePidConstants: function(jsonData) {
        this.queryById('pidConstantP').setValue(jsonData.p);
        this.queryById('pidConstantI').setValue(jsonData.i);
        this.queryById('pidConstantD').setValue(jsonData.d);
    },

    clearDebugOutput: function() {
        // console.log(message);
        // console.log(this.queryById('debugOutputContainer'));
        var dom = this.queryById('debugOutputContainer').el.dom;
        // console.log(dom);
        dom.innerHTML = "";
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
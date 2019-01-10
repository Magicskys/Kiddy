Vue.use(VueRouter);
const Info = { template: '<el-main>\n' +
        '                    <div style="margin-top: 20px">\n' +
        '                        <el-button v-bind:disabled="sqlinject" @click="toggleSelection()">取消选择</el-button>\n' +
        '                        <el-button v-bind:disabled="sqlinject" @click="toggleSqlInject()">SQL注入</el-button>\n' +
        '                        <el-button v-bind:disabled="sqlinject" @click="toggleXssInject()">XSS扫描</el-button>\n' +
        '                        <el-button v-bind:disabled="sqlinject" @click="toggleScan()">通用测试</el-button>\n' +
        '                    </div>\n' +
        '                    <el-table :data="tableData" @selection-change="handleSelectionChange" ref="multipleTable" tooltip-effect="dark">\n' +
        '                        <el-table-column type="selection" width="55">\n' +
        '                        </el-table-column>\n' +
        '                        <el-table-column prop="uid" label="id" width="55" v-if="show">\n' +
        '                        </el-table-column>\n' +
        '                        <el-table-column prop="method" label="类型" width="140">\n' +
        '                        </el-table-column>\n' +
        '                        <el-table-column prop="url" label="URL">\n' +
        '                        </el-table-column>\n' +
        '                        <el-table-column prop="cookies" label="Cookies">\n' +
        '                        </el-table-column>\n' +
        '                        <el-table-column prop="body" label="Body">\n' +
        '                        </el-table-column>\n' +
        '                    </el-table>\n' +
        '                </el-main>',
        data:function () {
            return{
                tableData:[],
                multipleSelection: [],
                sqlinject:true,
                show:false,
            }
        },
        created:function () {
            var self = this;
            this.toggleGetInfo();
        },
        methods: {
            toggleGetInfo(){
                var jsdata;
                $.ajax({
                    url: '/v1/info',
                    type: 'get',
                    data: {},
                    async : false,
                    dataType: 'json'
                }).then(function (res) {
                    if(res.code==200){
                        console.log(res);
                        //把从json获取的数据赋值给数组
                        jsdata= res.data;
                    }
                }).fail(function () {
                    console.log('失败');
                });
                this._data.tableData=jsdata;
            },
            toggleSelection() {
                this.$refs.multipleTable.clearSelection();
                this._data.sqlinject=true;
            },
            handleSelectionChange(val) {
                this._data.multipleSelection = val;
                if (this._data.multipleSelection.length==0) {
                    this._data.sqlinject=true;
                }else{
                    this._data.sqlinject=false;
                }
            },
            toggleSqlInject(){
                var idlist=[];
                for(i=0;i<this._data.multipleSelection.length;i++){
                    idlist.push(this._data.multipleSelection[i].uid)
                }
                if(request_scan(idlist,"sql")){
                    message_echo(this,'开始SQL注入，共注入测试'+this._data.multipleSelection.length+'条数据',"success");
                }else{
                    message_echo(this,'SQL注入准备失败，请检测网络状况',"error");
                }
            },
            toggleXssInject(){
                message_echo(this,'开始XSS注入，共注入测试'+this._data.multipleSelection.length+'条数据',"success");
            },
            toggleScan(){
                var idlist=[];
                for(i=0;i<this._data.multipleSelection.length;i++){
                    idlist.push(this._data.multipleSelection[i].uid)
                }
                if(request_scan(idlist,"general")) {
                    message_echo(this, '开始通用，共通用测试' + this._data.multipleSelection.length + '条数据', "success");
                }else{
                    message_echo(this, '通用测试准备失败，请检测网络状况', "error");
                }
            },
        },
};

const ScanResultSql={ template: '<el-main>\n' +
        '                <div style="margin-top: 20px">\n' +
        '                    <el-button v-bind:disabled="startbutton" @click="toggleSelection()">取消选择</el-button>\n' +
        '                    <el-button v-bind:disabled="startbutton" @click="toggleAction(\'start\')">开始</el-button>\n' +
        '                    <el-button v-bind:disabled="startbutton" @click="toggleAction(\'stop\')">暂停</el-button>\n' +
        '                    <el-button v-bind:disabled="startbutton" @click="toggleAction(\'kill\')">删除</el-button>\n' +
        '                </div>\n' +
        '                <el-table :data="tableData" style="width: 100%" :row-class-name="tableRowClassName" ref="multipleTable" tooltip-effect="dark" @selection-change="handleSelectionChange">\n' +
        '                    <el-table-column type="selection" width="55">\n' +
        '                    </el-table-column>\n' +
        '                    <el-table-column prop="taskId" label="id" width="55" v-if="show">\n' +
        '                    </el-table-column>\n' +
        '                    <el-table-column prop="datetime" label="日期" width="180">\n' +
        '                    </el-table-column>\n' +
        '                    <el-table-column prop="result2[0].url" label="URL" width="450">\n' +
        '                    </el-table-column>\n' +
        '                    <el-table-column prop="result2[0].method" label="method" width="180">\n' +
        '                    </el-table-column>\n' +
        '                    <el-table-column prop="result[0].status" label="状态">\n' +
        '                    </el-table-column>\n' +
        '                    <el-table-column fixed="right" label="操作" width="100">\n' +
        '                        <template slot-scope="scope">\n' +
        '                            <el-button @click="handleClick(scope,\'start\')" type="text" size="small">开始</el-button>\n' +
        '                            <el-button @click="handleClick(scope,\'kill\')" type="danger" size="small">删除</el-button>\n' +
        '                        </template>\n' +
        '                    </el-table-column>\n' +
        '                </el-table>\n' +
        '            </el-main>',
    data:function () {
        return{
            tableData:[],
            multipleSelection: [],
            startbutton:true,
            show:false,
        }
    },
    created:function () {
        var self = this;
        this.toggleGetInfo();
    },
    methods: {
        toggleGetInfo(){
            var jsdata;
            $.ajax({
                url: '/v1/result/sql',
                type: 'get',
                data: {},
                async : false,
                dataType: 'json'
            }).then(function (res) {
                if(res.code==200){
                    console.log(res);
                    //把从json获取的数据赋值给数组
                    jsdata= res.data;
                }
            }).fail(function () {
                console.log('失败');
            });
            this._data.tableData=jsdata;
        },
        toggleSelection() {
            this.$refs.multipleTable.clearSelection();
            this._data.sqlinject=true;
        },
        handleSelectionChange(val) {
            this._data.multipleSelection = val;
            if (this._data.multipleSelection.length==0) {
                this._data.startbutton=true;
            }else{
                this._data.startbutton=false;
            }
        },
        SelectMessage(action,index){
            switch (action) {
                case "start":
                    message_echo(this,'共开始'+this._data.multipleSelection.length+'个任务',"success");
                    break;
                case "stop":
                    message_echo(this, '共暂停'+this._data.multipleSelection.length+'条任务',"success");
                    break;
                case "kill":
                    message_echo(this,'共删除'+this._data.multipleSelection.length+'个任务',"success");
                    if(typeof(index)=="string" || typeof(index)=="number"){
                        this._data.tableData.splice(index, 1);
                    }else{
                        for(i=0;i<index.length;i++){
                            this._data.tableData.splice(i, 1);
                        }
                    }
                    break;
            }

        },
        SelectMessageError(action){
            switch (action) {
                case "start":
                    message_echo(this,'共开始'+this._data.multipleSelection.length+'个任务',"success");
                    break;
                case "stop":
                    message_echo(this, '共暂停'+this._data.multipleSelection.length+'条任务',"success");
                    break;
                case "kill":
                    message_echo(this,'共删除'+(this._data.multipleSelection.length || 1)+'个任务',"success");
                    break;
            }

        },
        toggleAction(action){
            var idlist=[];
            var idlistid=[];
            for(i=0;i<this._data.multipleSelection.length;i++){
                idlist.push(this._data.multipleSelection[i].taskId)
                idlistid.push(i)
            }
            if(request_action_schema('sql',idlist,action)){
                this.SelectMessage(action,idlistid);
            }else{
                this.SelectMessageError(action);
            }
        },
        tableRowClassName({row, rowIndex}) {
            console.log(row.result[0].status);
            if (row.result[0].status === "running") {
                return 'success-row';
            } else if (rowIndex === 3) {
                return 'warning-row';
            }
            return '';
        },
        handleClick(scope,action) {
            var idlist=[scope.row.taskId];
            if(request_action_sql(idlist,action)){
                this.SelectMessage(action,scope.$index);
            }else{
                this.SelectMessageError(action);
            }
        },
    },
};

const ScanResultGeneral={ template: '<el-main>\n' +
        '                <div style="margin-top: 20px">\n' +
        '                    <el-button v-bind:disabled="startbutton" @click="toggleSelection()">取消选择</el-button>\n' +
        '                    <el-button v-bind:disabled="startbutton" @click="toggleAction(\'kill\')">删除</el-button>\n' +
        '                </div>\n' +
        '                <el-table :data="tableData" style="width: 100%" :row-class-name="tableRowClassName" ref="multipleTable" tooltip-effect="dark" @selection-change="handleSelectionChange">\n' +
        '                    <el-table-column type="selection" width="55">\n' +
        '                    </el-table-column>\n' +
        '                    <el-table-column prop="taskId" label="id" width="55" v-if="show">\n' +
        '                    </el-table-column>\n' +
        '                    <el-table-column prop="typeof" label="类型" width="100">\n' +
        '                    </el-table-column>\n' +
        '                    <el-table-column prop="datetime" label="日期" width="180">\n' +
        '                    </el-table-column>\n' +
        '                    <el-table-column prop="host" label="HOST" width="450">\n' +
        '                    </el-table-column>\n' +
        '                    <el-table-column prop="status" label="状态" width="180">\n' +
        '                    </el-table-column>\n' +
        '                    <el-table-column prop="message" width="180" v-if="show">\n' +
        '                    </el-table-column>\n' +
        '                    <el-table-column label="结果" width="180">\n' +
        '                        <template slot-scope="scope">\n' +
        '                            <el-button @click="SelectInfo(scope)" type="text" size="small">查看</el-button>\n' +
        '                        </template>\n' +
        '                    </el-table-column>\n' +
        '                    <el-table-column label="操作" width="100">\n' +
        '                        <template slot-scope="scope">\n' +
        '                            <el-button @click="handleClick(scope,\'kill\')" type="danger" size="small">删除</el-button>\n' +
        '                        </template>\n' +
        '                    </el-table-column>\n' +
        '                </el-table>\n' +
        '            </el-main>',
    data:function () {
        return{
            tableData:[],
            multipleSelection: [],
            startbutton:true,
            show:false,
        }
    },
    created:function () {
        var self = this;
        this.toggleGetInfo();
    },
    methods: {
        toggleGetInfo(){
            var jsdata;
            $.ajax({
                url: '/v1/result/general',
                type: 'get',
                data: {},
                async : false,
                dataType: 'json'
            }).then(function (res) {
                if(res.code==200){
                    console.log(res);
                    //把从json获取的数据赋值给数组
                    jsdata= res.data;
                }
            }).fail(function () {
                console.log('失败');
            });
            this._data.tableData=jsdata;
        },
        toggleSelection() {
            this.$refs.multipleTable.clearSelection();
            this._data.sqlinject=true;
        },
        handleSelectionChange(val) {
            this._data.multipleSelection = val;
            if (this._data.multipleSelection.length==0) {
                this._data.startbutton=true;
            }else{
                this._data.startbutton=false;
            }
        },
        SelectInfo(scope){
            if(scope.row.message!="" && scope.row.typeof=="nmap"){
                var x2js = new X2JS();
                var jsonObj = x2js.xml_str2json(scope.row.message);
                this.$alert(jsonObj.nmaprun.scaninfo._services, '扫描结果', {
                    dangerouslyUseHTMLString: true
                });
            }else{
                this.$alert(scope.row.message, '扫描结果', {
                    dangerouslyUseHTMLString: true
                });
            }
        },
        SelectMessage(action,index){
            switch (action) {
                case "start":
                    message_echo(this,'共开始'+this._data.multipleSelection.length+'个任务',"success");
                    break;
                case "stop":
                    message_echo(this, '共暂停'+this._data.multipleSelection.length+'条任务',"success");
                    break;
                case "kill":
                    message_echo(this,'共删除'+(this._data.multipleSelection.length || 1)+'个任务',"success");
                    if(typeof(index)=="string" || typeof(index)=="number"){
                        this._data.tableData.splice(index, 1);
                    }else{
                        for(i=0;i<index.length;i++){
                            this._data.tableData.splice(i, 1);
                        }
                    }
                    break;
            }
        },
        SelectMessageError(action){
            switch (action) {
                case "start":
                    message_echo(this,'共开始'+this._data.multipleSelection.length+'个任务',"success");
                    break;
                case "stop":
                    message_echo(this, '共暂停'+this._data.multipleSelection.length+'条任务',"success");
                    break;
                case "kill":
                    message_echo(this,'共删除'+this._data.multipleSelection.length+'个任务',"success");
                    break;
            }

        },
        toggleAction(action){
            var idlist=[];
            var idlistid=[];
            for(i=0;i<this._data.multipleSelection.length;i++){
                idlist.push(this._data.multipleSelection[i].taskId);
                idlistid.push(i)
            }
            if(request_action_schema('general',idlist,action)){
                this.SelectMessage(action,idlistid);
            }else{
                this.SelectMessageError(action);
            }
        },
        tableRowClassName({row, rowIndex}) {
            console.log(row.status);
            if (row.status === "end") {
                return 'success-row';
            } else if (row.status === "start") {
                return 'warning-row';
            }
            return '';
        },
        handleClick(scope,action) {
            var idlist=[scope.row.taskId];
            if(request_action_schema('general',idlist,action)){
                this.SelectMessage(action,scope.$index);
            }else{
                this.SelectMessageError(action);
            }
        },
    },
};

const SQLMapSetting = {
    template: '<el-row type="flex" class="row-bg" justify="center">\n' +
        '                <el-col :span="12">\n' +
        '                    <el-form ref="form" :model="form" :rules="rules" label-width="140px">\n' +
        '                        <el-form-item label="SQLmap 本地路径" prop="sqlmap_localhost">\n' +
        '                            <el-input v-model="form.sqlmap_localhost" placeholder="~/sqlmap/"></el-input>\n' +
        '                        </el-form-item>\n' +
        '                        <el-form-item label="SQLmap API" prop="sqlmap_api">\n' +
        '                            <el-input v-model="form.sqlmap_api" placeholder="127.0.0.1:8775"></el-input>\n' +
        '                        </el-form-item>\n' +
        '                        <el-form-item label="启动SQL注入组件">\n' +
        '                            <el-switch v-model="form.start"></el-switch>\n' +
        '                        </el-form-item>\n' +
        '                        <el-form-item label="自动更新刷新结果">\n' +
        '                            <el-switch v-model="form.refresh"></el-switch>\n' +
        '                        </el-form-item>\n' +
        '                        <el-form-item label="delay">\n' +
        '                            <el-input-number v-model="form.region" :min="1" :max="20" label="delay"></el-input-number>\n' +
        '                        </el-form-item>\n' +
        '                        <el-row type="flex" class="row-bg" justify="space-around">\n' +
        '                            <el-col :span="12">\n' +
        '                                <el-form-item label="risk">\n' +
        '                                    <el-input-number v-model="form.risk" :min="1" :max="3" label="risk"></el-input-number>\n' +
        '                                </el-form-item>\n' +
        '                            </el-col>\n' +
        '                            <el-col :span="12">\n' +
        '                                <el-form-item label="level">\n' +
        '                                    <el-input-number v-model="form.level" :min="1" :max="5" label="level"></el-input-number>\n' +
        '                                </el-form-item>\n' +
        '                            </el-col>\n' +
        '                        </el-row>\n' +
        '                        <el-form-item label="随机User-Agent">\n' +
        '                            <el-switch v-model="form.user_agent"></el-switch>\n' +
        '                        </el-form-item>\n' +
        '                        <el-form-item label="TECH">\n' +
        '                            <el-checkbox-group v-model="form.tech">\n' +
        '                                <el-checkbox label="B" name="type"></el-checkbox>\n' +
        '                                <el-checkbox label="E" name="type"></el-checkbox>\n' +
        '                                <el-checkbox label="U" name="type"></el-checkbox>\n' +
        '                                <el-checkbox label="S" name="type"></el-checkbox>\n' +
        '                                <el-checkbox label="T" name="type"></el-checkbox>\n' +
        '                            </el-checkbox-group>\n' +
        '                        </el-form-item>\n' +
        '                        <el-form-item>\n' +
        '                            <el-button type="primary" @click="submitForm(\'form\')">立即创建</el-button>\n' +
        '                            <el-button @click="resetForm(\'form\')">取消</el-button>\n' +
        '                        </el-form-item>\n' +
        '                    </el-form>\n' +
        '                </el-col>\n' +
        '            </el-row>',
    data:function () {
        return{
            form: {
                sqlmap_localhost:'',
                sqlmap_api: '',
                region: '',
                level: '',
                risk: '',
                start:false,
                refresh: false,
                tech: [],
                user_agent:false,
            },
            rules:{
                sqlmap_localhost:[
                    { required: true, message: '请输入SQLMAP 本地路径', trigger: 'blur' },
                ],
                sqlmap_api:[
                    { required: true, message: '请输入SQLMAPAPI 地址', trigger: 'blur' },
                ]
            },
        }
    },
    methods: {
        submitForm(formName){
            this.$refs[formName].validate((valid) => {
                var self=this;
                if (valid) {
                    var data={'sqlmap_localhost':this.$refs.form.model.sqlmap_localhost,
                        'sqlmap_api':this.$refs.form.model.sqlmap_api,
                        'region':this.$refs.form.model.region,
                        'level':this.$refs.form.model.level,
                        'risk':this.$refs.form.model.risk,
                        'start':this.$refs.form.model.start,
                        'refresh':this.$refs.form.model.refresh,
                        'tech[]':this.$refs.form.model.tech.join(""),
                        'user_agent':this.$refs.form.model.user_agent};
                    $.ajax({
                        url: '/v1/setting/sql',
                        type: 'POST',
                        data: data,
                        dataType: 'JSON',
                        ContentType: "application/json",
                        async : false,
                    }).then(function (res) {
                        if(res.code==200){
                            console.log(res);
                            message_echo(self,"更新数据成功","success");
                        }
                        if(res.code==400){
                            message_echo(self,"本地路径或本地路径下文件不存在","error");
                        }
                    }).fail(function () {
                        message_echo(self,"提交数据失败","error");
                        console.log('失败');
                    });
                } else {
                    console.log('error submit!!');
                    return false;
                }
            });
        },
        resetForm(formName) {
            this.$refs[formName].resetFields();
        }
    },
    created(){
        var self = this;
        $.ajax({
            url: '/v1/setting/sql',
            type: 'get',
            data: {},
            async : false,
            dataType: 'json'
        }).then(function (res) {
            if(res.code==200){
                console.log(res);
                self.form.refresh=res.data.sqlmap.auto_refresh;
                self.form.sqlmap_api=res.data.sqlmap.sqlmap_api;
                self.form.sqlmap_localhost=res.data.sqlmap.sqlmap_location;
                self.form.level=res.data.sqlmap.Options.level;
                self.form.start=res.data.sqlmap.start;
                self.form.region=res.data.sqlmap.Options.delay;
                self.form.user_agent=res.data.sqlmap.Options.randomagent;
                self.form.tech=res.data.sqlmap.Options.tech.split("");

            }
        }).fail(function () {
            console.log('失败');
        });
    },
};

const GenaralSetting={
    template: '<el-row type="flex" class="row-bg" justify="center">\n' +
        '                <el-col :span="12">\n' +
        '                    <el-form ref="form" :model="form" :rules="rules" label-width="80px">\n' +
        '                        <el-form-item label="插件设置">\n' +
        '                            <el-transfer filterable :filter-method="form.filterMethod" filter-placeholder="请输入插件名称" :titles="[\'插件列表\', \'已选中插件\']" v-model="form.value" :data="form.data">\n' +
        '                            </el-transfer>\n' +
        '                        </el-form-item> ' +
        '                        <el-form-item label="端口扫描">\n' +
        '                           <el-switch v-model="form.ports" active-text="快速扫描" active-color="#13ce66" inactive-color="#ff4949" inactive-text="自定义扫描">\n' +
        '                           </el-switch>\n' +
        '                        </el-form-item> ' +
        '                        <el-form-item label="自定义端口范围" label-width="unset">\n' +
        '                           <el-slider v-model="form.portrange" style="margin-left:140px;" v-bind:disabled="form.ports" :min="1" :max="65535" :step="1" range show-input>\n' +
        '                           </el-slider>\n'+
        '                        </el-form-item> ' +
        '                      <el-form-item label="扫描类型">\n'+
        '                        <el-radio-group v-model="form.portschema">\n'+
        '                           <el-row :gutter="20">\n'+
        '                               <el-col :span="8"><el-radio label="sP">sP ping扫描</el-radio></el-col>\n'+
        '                               <el-col :span="8"><el-radio label="sA">sA 发送tcp的ack包</el-radio></el-col>\n'+
        '                               <el-col :span="8"><el-radio label="sS">sS 半开放扫描（非3次握手的tcp扫描）</el-radio></el-col>\n'+
        '                               <el-col :span="8"><el-radio label="sT">sT 3次握手方式tcp的扫描</el-radio></el-col>\n'+
        '                               <el-col :span="8"><el-radio label="sU">sU udp端口的扫描</el-radio></el-col>\n'+
        '                               <el-col :span="8"><el-radio label="sF">sF 发送一个FIN标志的数据包</el-radio></el-col>\n'+
        '                               <el-col :span="8"><el-radio label="sX"></el-radio></el-col>\n'+
        '                               <el-col :span="8"><el-radio label="sN"></el-radio></el-col>\n'+
        '                               <el-col :span="8"><el-radio label="sW">sW 窗口扫描</el-radio></el-col>\n'+
        '                               <el-col :span="8"><el-radio label="sV">sV 版本检测</el-radio></el-col>\n'+
        '                           </el-row>\n'+
        '                         </el-radio-group>\n'+
        '                      </el-form-item>\n'+
        '                        <el-form-item>\n' +
        '                            <el-button type="primary" @click="submitForm(\'form\')">立即创建</el-button>\n' +
        '                            <el-button @click="resetForm(\'form\')">重置</el-button>\n' +
        '                        </el-form-item>\n' +
        '                    </el-form>\n' +
        '                </el-col>\n' +
        '            </el-row>',
    data:function () {
        return{
            form: {
                data: [],
                value: [],
                ports:true,
                portrange:65535,
                portschema:'',
                filterMethod(query, item) {
                    return item.pinyin.indexOf(query) > -1;
                }
            },
            rules:{
            },
        }
    },
    methods: {
        submitForm(formName){
            this.$refs[formName].validate((valid) => {
                var self=this;
                if (valid) {
                    var data={'porttype':this.$refs.form.model.ports,
                        'portrange':this.$refs.form.model.portrange,
                        'plugin':this.$refs.form.model.value,
                        'portschema':this.$refs.form.model.portschema};
                    $.ajax({
                        url: '/v1/setting/general',
                        type: 'POST',
                        data: data,
                        dataType: 'JSON',
                        ContentType: "application/json",
                        async : false,
                    }).then(function (res) {
                        if(res.code==200){
                            console.log(res);
                            message_echo(self,"更新数据成功","success")
                        }
                    }).fail(function () {
                        message_echo(self,"提交数据失败","error")
                        console.log('失败');
                    });
                } else {
                    console.log('error submit!!');
                    return false;
                }
            });
        },
        resetForm(formName) {
            this.$refs[formName].resetFields();
        }
    },
    created(){
        var self = this;
        $.ajax({
            url: '/v1/setting/general',
            type: 'GET',
            async : false,
            dataType: 'json'
        }).then(function (res) {
            if(res.code==200){
                console.log(res);
                self.form.ports=res.data.general.port_scan;
                self.form.portrange=res.data.general.port_range;
                self.form.portschema=res.data.general.portschema;
                for(var key in res.plugins){
                    self.form.data.push({
                        label: res.plugins[key].title,
                        key: res.plugins[key].pinyin,
                        pinyin: res.plugins[key].pinyin
                    });
                }
                for(var key in res.data.general.plugin){
                    self.form.value.push(res.data.general.plugin[key]);
                }
            }
        }).fail(function () {
            message_echo(self,"网络连接失败","error")
        });
    },
};

const PocList={ template: '<el-main>\n' +
        '                <div style="margin-top: 20px">\n' +
        '                    <el-button v-bind:disabled="startbutton" @click="toggleSelection()">取消选择</el-button>\n' +
        '                    <el-button v-bind:disabled="startbutton" @click="toggleAction(\'kill\')">删除</el-button>\n' +
        '                </div>\n' +
        '                <el-table :data="tableData" style="width: 100%" ref="multipleTable" tooltip-effect="dark" @selection-change="handleSelectionChange">\n' +
        '                    <el-table-column type="selection" :selectable="checkboxT" width="55">\n' +
        '                    </el-table-column>\n' +
        '                    <el-table-column prop="id" label="id" width="55" v-if="show">\n' +
        '                    </el-table-column>\n' +
        '                    <el-table-column prop="classification" label="类型" width="100">\n' +
        '                    </el-table-column>\n' +
        '                    <el-table-column prop="danger" label="危险评级" width="100">\n' +
        '                    </el-table-column>\n' +
        '                    <el-table-column prop="pinyin" width="100" v-if="show">\n' +
        '                    </el-table-column>\n' +
        '                    <el-table-column prop="title" label="插件名称" width="180">\n' +
        '                    </el-table-column>\n' +
        '                    <el-table-column label="操作" width="100">\n' +
        '                        <template slot-scope="scope">\n' +
        '                            <el-button @click="handleClick(scope,\'kill\')" type="danger" size="small">删除</el-button>\n' +
        '                        </template>\n' +
        '                    </el-table-column>\n' +
        '                </el-table>\n' +
        '            </el-main>',
    data:function () {
        return{
            tableData:[],
            multipleSelection: [],
            startbutton:true,
            show:false,
        }
    },
    created:function () {
        var self = this;
        this.toggleGetInfo();
    },
    methods: {
        toggleGetInfo(){
            var jsdata;
            $.ajax({
                url: '/v1/poc/list',
                type: 'get',
                data: {},
                async : false,
                dataType: 'json'
            }).then(function (res) {
                if(res.code==200){
                    console.log(res);
                    //把从json获取的数据赋值给数组
                    jsdata= res.plugins;
                }
            }).fail(function () {
                console.log('失败');
            });
            this._data.tableData=jsdata;
        },
        toggleSelection() {
            this.$refs.multipleTable.clearSelection();
            this._data.sqlinject=true;
        },
        handleSelectionChange(val) {
            this._data.multipleSelection = val;
            if (this._data.multipleSelection.length==0) {
                this._data.startbutton=true;
            }else{
                this._data.startbutton=false;
            }
        },
        SelectMessage(action,index){
            switch (action) {
                case "kill":
                    message_echo(this,'共删除'+(this._data.multipleSelection.length || 1)+'个任务',"success");
                    if(typeof(index)=="string" || typeof(index)=="number"){
                        this._data.tableData.splice(index, 1);
                    }else{
                        for(i=0;i<index.length;i++){
                            this._data.tableData.splice(i, 1);
                        }
                    }
                    break;
            }
        },
        SelectMessageError(action){
            switch (action) {
                case "kill":
                    message_echo(this,'共删除'+(this._data.multipleSelection.length || 1)+'个任务',"success");
                    break;
            }

        },
        toggleAction(action){
            var idlist=[];
            var idlistid=[];
            for(i=0;i<this._data.multipleSelection.length;i++){
                idlist.push(this._data.multipleSelection[i].id);
                idlistid.push(i)
            }
            if(request_action_schema('poc',idlist,action)){
                this.SelectMessage(action,idlistid);
            }else{
                this.SelectMessageError(action);
            }
        },
        checkboxT(row, index) {
            if (row.pinyin in {"portscan":null,"xssinject":null,"sqlinject":null}){
                return 0;
            }
            return 1;
        },
        handleClick(scope,action) {
            var idlist=[scope.row.id];
            if(request_action_schema('poc',idlist,action)){
                this.SelectMessage(action,scope.$index);
            }else{
                this.SelectMessageError(action);
            }
        },
    },
};

const PocAdd = {
    template: '<el-row type="flex" class="row-bg" justify="center">\n' +
        '                <el-col :span="12">\n' +
        '                    <el-form ref="form" :model="form" :rules="rules" label-width="140px">\n' +
        '                        <el-form-item label="POC标题" prop="title">\n' +
        '                            <el-input v-model="form.title" placeholder="XXX目录遍历"></el-input>\n' +
        '                        </el-form-item>\n' +
        '                        <el-form-item label="POC分类" prop="classification">\n' +
        '                            <el-input v-model="form.classification" placeholder="CMS"></el-input>\n' +
        '                        </el-form-item>\n' +
        '                        <el-form-item label="危险程度" prop="danger" :error="errorMsg">\n' +
        '                           <el-select v-model="form.danger" placeholder="请选择危险评级">\n' +
        '                               <el-option label="低" value="low"></el-option>\n' +
        '                               <el-option label="中" value="medium"></el-option>\n' +
        '                               <el-option label="高" value="high"></el-option>\n' +
        '                           </el-select>\n' +
        '                        </el-form-item>\n'+
        '                        <el-form-item label="POC" prop="poc">\n' +
        '                        <el-button type="text" @click="open">点击打开POC编写事例</el-button>\n'+
        '                          <el-input type="textarea" v-model="form.poc" rows="20"></el-input>\n' +
        '                        </el-form-item>\n'+
        '                        <el-form-item>\n' +
        '                            <el-button type="primary" @click="submitForm(\'form\')">立即添加</el-button>\n' +
        '                            <el-button @click="resetForm(\'form\')">重置</el-button>\n' +
        '                        </el-form-item>\n' +
        '                    </el-form>\n' +
        '                </el-col>\n' +
        '            </el-row>',
    data:function () {
        return{
            form: {
                title:'',
                danger: '',
                classification:'',
                poc: '',
            },
            errorMsg:"",
            rules:{
                title:[
                    { required: true, message: '请输入POC标题', trigger: 'blur' },
                ],
                classification:[
                    { required: true, message: '请输入POC类别', trigger: 'blur' },
                ],
                danger:[
                    { required: true, message: '请选择危险评级', trigger: 'blur' },
                ],
                poc:[
                    { required: true, message: '请输入POC', trigger: 'blur' },
                ],
            },
        }
    },
    methods: {
        submitForm(formName){
            this.$refs[formName].validate((valid) => {
                var self=this;
                // var formName=formName;
                if (valid) {
                    var data={'title':this.$refs.form.model.title,
                        'danger':this.$refs.form.model.danger,
                        'classification':this.$refs.form.model.classification,
                        'poc':this.$refs.form.model.poc,
                        'pinyin':this.$refs.form.model.pinyin,
                        };
                    this.errorMsg= '';
                    $.ajax({
                        url: '/v1/setting/poc',
                        type: 'POST',
                        data: data,
                        dataType: 'JSON',
                        ContentType: "application/json",
                        async : false,
                    }).then(function (res) {
                        if(res.code==200){
                            console.log(res);
                            message_echo(self,"添加POC成功","success");
                            self.$refs[formName].resetFields();
                        }
                        if(res.code==202){
                            self.errorMsg="输入错误请重新输入";
                            message_echo(self,"POC分类 字段错误请重新输入","error");
                        }
                    }).fail(function () {
                        message_echo(self,"添加POC失败","error");
                        console.log('失败');
                    });
                } else {
                    console.log('error submit!!');
                    return false;
                }
            });
        },
        resetForm(formName) {
            this.$refs[formName].resetFields();
        },
        open() {
            this.$alert('<p>import requests</p>' +
                '<p>requests.get("<strong style="color: blue;">{#url#}</strong>")</p>' +
                '<p>print("<strong style="color: blue;">[***]</strong>存在XXX漏洞")</p>'+
                '<p>其中<strong style="color: blue;">{#url#}</strong>为URL标识位</p>'+
                '<p>其中<strong style="color: blue;">[***]</strong>为漏洞发现标识位</p>', 'POC事例', {
                dangerouslyUseHTMLString: true,
            });
        },
    },
};


const routes = [
    { path: '/info', component: Info },
    { path: '/setting/sql', component: SQLMapSetting },
    { path: '/setting/general', component: GenaralSetting },
    { path: '/result/sql', component:ScanResultSql },
    { path: '/result/general', component:ScanResultGeneral },
    { path: '/poc/list', component:PocList },
    { path: '/poc/add', component:PocAdd },
];

const router = new VueRouter({
    routes:routes
});
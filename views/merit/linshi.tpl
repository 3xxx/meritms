<!-- $("#customerName").click(getConfirm);

    function getConfirm(){
      
      // 删除，保证每次都是最新的数据
      $(".on_changes li").remove();
      
      // 加载客户数据
      $.ajax({
          url:'public/getConfirm',
        type:'GET',
        data:"id="+"${userId}",
        success:function(data){
          //alert(data);
          var arr = ""+data+"";
            var dataObj = eval("("+arr+")");
            
             $.each(dataObj.rows,function(idx,item){
               //组装数据
               //alert(item);
               var li = "<li onclick='get(&apos;"+item[0]+"&apos;,&apos;"+item[1]+"&apos;)' onmouseover='this.style.backgroundColor=\"#ffff66\";'onmouseout='this.style.backgroundColor=\"#fff\";'>"+item[1]+"</li>";
               //alert(li);
               $(".on_changes").append(li);
            });
        }
      });
      
      // 控制下拉框显示
      var display =$('.on_changes');
      if(display.is(':hidden')){//如果node是隐藏的则显示node元素，否则隐藏
        $(".on_changes").show();
      }else{
        $(".on_changes").hide();
      }
    }
    
    function get(data1,data2){
      //alert(data1);// 客户Id
      //alert(data2);// 客户名称
    
      $("#customerName").val(data2);
    
      $(".on_changes").hide();
    }


    css
    .on_changes{width:232px; position:absolute; top:61px; list-style:none; background:#FFF; border:1px solid #000; display:none; padding:3px;}
.on_changes li{margin:0px;padding:6px;font-size: 14px;}
.on_changes li.active{ background:#CEE7FF;} -->

<!-- <input type="text" class="salesInfo" id="customerName"/>
<ul class="on_changes"> -->
<!-- <li onclick="get('4df5','sdf')">sdf</li> -->
<!-- </ul>


  var flag3 = 0;
  function create_table2(){
   // var num = document.getElementById("num2").value;
   // if(num == null ){
    // num = 1;
   // }
   var s1 = document.getElementById("d2");
   var vTable = document.createElement("Table");
   // for(i=0; i<num; i++){
    vTr = vTable.insertRow();
    vTd = vTr.insertCell();
    vTd.innerHTML="三级目录"+(flag3+1)+"：<input type=\"text\" name=\"category3\" id=\"category3\" />";
   // }//type=\"file\"
   s1.appendChild(vTable);
   flag3++;
  } 

    var flag4 = 0; 
   function create_table3(){
    var rowId = "row" + flag4;
   // var num = document.getElementById("num3").value;
   // if(num == null ){
    // num = 1;
   // }
   var s1 = document.getElementById("d3");
   var vTable = document.createElement("Table");
   // for(i=0; i<num; i++){
    vTr = vTable.insertRow();
    vTd = vTr.insertCell();
    vTd.innerHTML="四级目录"+(flag4+1)+"：<input type=\"text\" name=\"category4\" id=\"category4\" /> <span>选择类别</span> <input type=\"radio\" id='txtIndex" +flag4+ "' name='radiobutton"+flag4+"' checked='true' value='Attachment'/><label for='p_man'>附件模式</label><input type=\"radio\"  id='txtIndex" +flag4+ "' name='radiobutton"+flag4+"' value='Fdiary'/><label for='p_man'>图文模式</label>";
   // }//type=\"file\"
   s1.appendChild(vTable);
   flag4++;
  } 



  <script type="text/javascript">
$(document).ready(function(){
  $("input").focus(function(){
    $("input").css("background-color","#FFFFCC");
  });
  $("input").blur(function(){
    $("input").css("background-color","#D6D6FF");
  });
  $("#btn1").click(function(){
    $("input").focus();
  });  
  $("#btn2").click(function(){
    $("input").blur();
  }); 
});
</script>

var objArr = document.getElementsByName('radiogroup1');
         
            for (var i = 0; i < objArr.length; i++)
            {
                alert(objArr[i].id);
            }



function noEditFormatter(cellValue, options, rowObject) {
 if (cellValue == 'test')
  jQuery("#grid").jqGrid('setCell', options.rowId, 'ColName', '', 'not-editable-cell');
 return cellValue;
}

onLoadSuccess:function(data){
  setTimeout(function(){
  $('#table')[0].rows[1].cells[2].innerHTML='';
  },100)
}


$(function () {
    var $table = $('#table').bootstrapTable({
        idField: 'name',
        url: '/gh/get/response.json/wenzhixin/bootstrap-table/tree/master/docs/data/data1/',
        columns: [{
            field: 'name',
            title: 'Name'
        }, {
            field: 'stargazers_count',
            title: 'Stars',
            editable: {
                type: 'text'
            }
        }, {
            field: 'forks_count',
            title: 'Forks',
            editable: {
                type: 'text'
            }
        }, {
            field: 'description',
            title: 'Description',
            editable: {
                type: 'textarea'
            }
        }]
    });
    $table.on('editable-save.bs.table', function (e, field, row, old, $el) {
        var $els = $table.find('.editable'),
            next = $els.index($el) + 1;

            if (next >= $els.length) {
                return;
            }
            $els.eq(next).editable('show');
    });
});


$(function () {
    $('#table').bootstrapTable({
        idField: 'name',
        pagination: true,
        search: true,
        url: '/gh/get/response.json/wenzhixin/bootstrap-table/tree/master/docs/data/data1/',
        columns: [{
            field: 'name',
            title: 'Name'
        }, {
            field: 'stargazers_count',
            title: 'Stars'
        }, {
            field: 'forks_count',
            title: 'Forks'
        }, {
            field: 'description',
            title: 'Description'
        }],
        onPostBody: function () {
            $('#table').editableTableWidget({editor: $('<textarea>')});
        }
    });
});
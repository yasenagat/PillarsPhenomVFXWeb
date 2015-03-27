<!doctype html>
<html>
	<head>
		<title>Project</title>
		<meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
		<link href="style/css.css" rel="stylesheet" type="text/css">
		<link href="style/project.css" rel="stylesheet" type="text/css">
		<script language="JavaScript" src="jslib/laydate/laydate.js"></script>
		<script>
		function datatimes(str){
			var strlength = str.indexOf(" ");
			return str.substring(0,strlength);
		}
		</script>
	</head>

	<body>
		<div class="menu">
			<div class="menuall">
			<div class="logo"></div>
				<!--<ul>
				  <li><a href="project_list">project</a></li>
				  <li><a href="#" class="editoralpage">editoral</a></li>
				  <li><a href="#">post</a></li>
				  <li><a href="#">on-set</a></li>
				  <li><a href="#">playback</a></li>
				</ul>-->
			</div>
			<div class="menuab">
						 <span><img src="img/menu.png"></span>
					<div class="menu_tabs">
								<ul>
										<li>
											<a href="login.html">
											退出
											</a>
							</li>
								  </ul>
					</div>
				</div>
		</div>
		<div class="all">
			<div class="cont">
				<div class="search">
					<span class="spsearch">
						<input type="text" class="inptext" placeholder="搜索">
						<input type="hidden" id="searchflag" value="">
						<input type="button" value="" class="inpsearch">
					</span>
					<input type="button" value="添加项目" class="inpbutton">
				</div>
				<div class="pro">

					<!--从这里开始遍历-->
			{{with .}}
				{{range .}}
					<div class="protab">
						<table width="100%" border="0" cellspacing="0" cellpadding="0" class="tab{{.ProjectCode}}">
							<tbody>
								<tr>
									<td rowspan="3" width="150" height="200">
										<script type="text/javascript">
											document.write( "<img width='120' height='160' src={{.Picture}}>" );
										</script>
									</td>
									<td width="15%" align="right">项目名：
										<input type="hidden" class="ProjectCode_" value="{{.ProjectCode}}">
									</td>
									<td class="ProjectName_">
										{{.ProjectName}}
									</td>
									<td width="15%" align="right">负责人：</td>
									<td class="ProjectLeader_">
										{{.ProjectLeader}}
									</td>
								</tr>
								<tr>
									<td align="right">开始时间：</td>
									<td class="StartDatetime_">
									<script type="text/javascript">
						document.write(datatimes("<div>"+{{.StartDatetime}}+"</div>"));
									</script>
									</td>
									<td align="right">结束时间：</td>
									<td class="EndDatetime_">
									<script type="text/javascript">				
						document.write("<div>"+datatimes({{.EndDatetime}})+"</div>");
									</script>
									</td>
								</tr>
								<tr>
									<td align="right">项目类型：</td>
									<td class="ProjectType_">
										{{.ProjectType}}
									</td>
									<td align="right">备注：</td>
									<td class="ProjectDetail_">
										{{.ProjectDetail}}
									</td>
								</tr>
							</tbody>
						</table>
						<div class="cz">操作</div>
						<div class="del">删除</div>
					</div>

				{{end}}
			{{end}}
					<!--结束遍历-->

				</div>
			</div>
		</div>
		<div class="outer"></div>
		<div class="formdiv">
			<form id="projectForm" method="post">
				<table width="100%" height="400" border="0" cellspacing="0" cellpadding="0" class="formtable">
					<tr>
						<td width="30%" align="right">项目名：</td>
						<td width="70%">
							<input type="text" name="ProjectName" class="ProjectName">
							<input type="hidden" class="ProjectCode">
						</td>
					</tr>
					<tr>
						<td align="right">项目类型：</td>
						<td><input type="text" name="ProjectType" class="ProjectType"></td>
					</tr>
					<tr>
						<td align="right">开始时间：</td>
						<td><input type="text" name="StartDatetime" class="StartDatetime" onclick="laydate()"></td>
					</tr>
					<tr>
						<td align="right">结束时间：</td>
						<td><input type="text" name="EndDatetime" class="EndDatetime" onclick="laydate()"></td>
					</tr>
					<tr>
						<td align="right">负责人：</td>
						<td><input type="text" name="ProjectLeader" class="ProjectLeader"></td>
					</tr>
					<tr>
						<td align="right">备注：</td>
						<td><input type="text" name="ProjectDetail" class="ProjectDetail"></td>
					</tr>
					<tr>
						<td colspan="2" align="center">
							<input type="button" class="again" value="重置">
							<input type="button" class="submit" value="提交">
						</td>
					</tr>
				</table>
			</form>
		</div>

		<script data-main="js/project" language="JavaScript" defer async="true" src="jslib/require.js"></script>
	</body>
</html>


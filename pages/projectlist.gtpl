<!doctype html>
<html>
	<head>
		<title>项目列表</title>
		<meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
		<link rel="stylesheet" type="text/css" href="style/userManage.css"/>
		<script type="text/javascript" src="/jslib/jquery-1.4.3.js"></script>
	</head>

	<body>
		<div id="wrap">
			<div id="top_content">
				<div id="header">
					<div id="rightheader">
						<p>
							2015/01/05
							<br />
						</p>
					</div>
					<div id="topheader">
						<h1 id="title">
							<a href="#">项目管理</a>
						</h1>
					</div>
					<div id="navigation">
					</div>
				</div>
				<div id="content">

					<h1>
						项目列表
					</h1>
					<table class="table">
						<tr class="table_header">
							<td>
								项目名称
							</td>
							<td>
								项目描述
							</td>
							<td>
								项目负责人
							</td>
							<td>
								操作
							</td>
						</tr>

				{{with .}}
					{{range .}}
						<tr class="row{{.Status}}">
							<td>
								{{.ProjectName}}
							</td>
							<td>
								{{.ProjectDescribe}}
							</td>
							<td>
								{{.ProjectLeader}}
							</td>
							<td>
								<a href="javascript:void(0)" onclick="javascript:deleteproject('{{.ProjectCode}}')">删除</a>&nbsp;<a href="javascript:void(0)" onclick="javascript:updateproject('{{.ProjectCode}}')">修改</a>
							</td>
						</tr>
					{{end}}
				{{end}}

					</table>
					<p>
						<input type="button" class="button" value="新增项目" onclick="location='addUser.html'"/>
					</p>
				</div>
			</div>
			<div id="footer">
				<div id="footer_bg">
					chengxuezhi@git.com
				</div>
			</div>
		</div>

		<script type="text/javascript" src="/js/projectlist.js"></script>
	</body>
</html>

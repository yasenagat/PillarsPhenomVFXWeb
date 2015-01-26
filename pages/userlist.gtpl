<!doctype html>
<html>
	<head>
		<title>用户列表</title>
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
							<a href="#">用户管理</a>
						</h1>
					</div>
					<div id="navigation">
					</div>
				</div>
				<div id="content">

					<h1>
						用户列表
					</h1>
					<table class="table">
						<tr class="table_header">
							<td>
								Email
							</td>
							<td>
								用户名
							</td>
							<td>
								电话
							</td>
							<td>
								分组
							</td>
							<td>
								文件路径
							</td>
							<td>
								创建日期
							</td>
							<td>
								修改日期
							</td>
							<td>
								操作
							</td>
						</tr>

				{{with .}}
					{{range .}}
						<tr class="row{{.Status}}">
							<td>
								{{.Email}}
							</td>
							<td>
								{{.DisplayName}}
							</td>
							<td>
								{{.Phone}}
							</td>
							<td>
								{{.UserAuthority}}
							</td>
							<td>
								{{.FilePath}}
							</td>
							<td>
								{{.InsertDatetime}}
							</td>
							<td>
								{{.UpdateDatetime}}
							</td>
							<td>
								<a href="javascript:void(0)" onclick="javascript:deleteuser('{{.Email}}')">删除</a>&nbsp;<a href="javascript:void(0)" onclick="javascript:updateuser('{{.Email}}')">修改</a>
							</td>
						</tr>
					{{end}}
				{{end}}

					</table>
					<p>
						<input type="button" class="button" value="新增用户" onclick="location='addUser.html'"/>
					</p>
				</div>
			</div>
			<div id="footer">
				<div id="footer_bg">
					test@mail.com
				</div>
			</div>
		</div>

		<script type="text/javascript" src="/js/userlist.js"></script>
	</body>
</html>

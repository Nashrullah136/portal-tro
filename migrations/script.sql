/****** Object:  Table [dbo].[audits]    Script Date: 12/07/2023 16:09:33 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[audits](
	[date_time] [datetime] NOT NULL,
	[id] [bigint] IDENTITY(1,1) NOT NULL,
	[username] [varchar](32) NOT NULL,
	[entity] [varchar](32) NULL,
	[entity_id] [varchar](50) NULL,
	[action] [varchar](144) NULL,
	[data_before] [nvarchar](512) NULL,
	[data_after] [nvarchar](512) NULL,
 CONSTRAINT [audits_pk] PRIMARY KEY CLUSTERED 
(
	[id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY]
GO
/****** Object:  Table [dbo].[roles]    Script Date: 12/07/2023 16:09:33 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[roles](
	[id] [bigint] IDENTITY(1,1) NOT NULL,
	[role_name] [varchar](32) NOT NULL,
 CONSTRAINT [roles_pk] PRIMARY KEY CLUSTERED 
(
	[id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY]
GO
/****** Object:  Table [dbo].[users]    Script Date: 12/07/2023 16:09:33 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[users](
	[id] [bigint] IDENTITY(1,1) NOT NULL,
	[username] [varchar](32) NOT NULL,
	[password] [varchar](255) NOT NULL,
	[role_id] [bigint] NOT NULL,
	[created_at] [datetime] NOT NULL,
	[created_by] [varchar](32) NOT NULL,
	[updated_at] [datetime] NOT NULL,
	[updated_by] [varchar](32) NOT NULL,
	[name] [varchar](64) NULL,
 CONSTRAINT [users_pk] PRIMARY KEY CLUSTERED 
(
	[id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY]
GO

SET IDENTITY_INSERT [dbo].[roles] ON
INSERT [dbo].[roles] ([id], [role_name]) VALUES (1, N'admin')
INSERT [dbo].[roles] ([id], [role_name]) VALUES (2, N'user')
SET IDENTITY_INSERT [dbo].[roles] OFF
GO

SET IDENTITY_INSERT [dbo].[users] ON
INSERT [dbo].[users] ([id], [username], [password], [role_id], [created_at], [created_by], [updated_at], [updated_by], [name]) VALUES (1, N'admin', N'$2a$10$HogkQLhuLXdCBrJwS98c9ezB/mhAJn/k7Ru7HKwohVFKttPXeKy6a', 1, CAST(N'2023-06-14T13:46:02.000' AS DateTime), N'admin', CAST(N'2023-07-09T20:08:13.113' AS DateTime), N'admin', N'ADMIN')
SET IDENTITY_INSERT [dbo].[users] OFF
GO

SET ANSI_PADDING ON
GO
/****** Object:  Index [username_unique]    Script Date: 12/07/2023 16:09:33 ******/
ALTER TABLE [dbo].[users] ADD  CONSTRAINT [username_unique] UNIQUE NONCLUSTERED 
(
	[username] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, SORT_IN_TEMPDB = OFF, IGNORE_DUP_KEY = OFF, ONLINE = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
GO
ALTER TABLE [dbo].[audits] ADD  DEFAULT (getdate()) FOR [date_time]
GO
ALTER TABLE [dbo].[users] ADD  DEFAULT (getdate()) FOR [created_at]
GO
ALTER TABLE [dbo].[users] ADD  DEFAULT (getdate()) FOR [updated_at]
GO
ALTER TABLE [dbo].[users]  WITH CHECK ADD  CONSTRAINT [role_id_users] FOREIGN KEY([role_id])
REFERENCES [dbo].[roles] ([id])
GO
ALTER TABLE [dbo].[users] CHECK CONSTRAINT [role_id_users]
GO

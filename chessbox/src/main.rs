use bevy::{math::Vec3Swizzles, prelude::*};

mod component;
mod resource;


fn main() {
    // add_startup_system 启动生命周期时只运行一次
    // add_system 每帧都会调用方法
    App::new()
        .add_state::<resource::GameState>()
        .insert_resource(ClearColor(Color::rgb(255., 255., 255.)))
        .add_plugins(DefaultPlugins.set(WindowPlugin {
            primary_window: Some(Window {
                title: "中国象棋".into(),
                resolution: (1060., 808.).into(),
                ..Window::default()
            }),
            ..WindowPlugin::default()
        }))
        .add_systems(Startup, setup_system)
        // 启动 esc 键退出程序
        .add_systems(Update, bevy::window::close_on_esc)
        .run();
}

fn setup_system(
    mut commands: Commands,
    asset_server: Res<AssetServer>,
    mut texture_atlases: ResMut<Assets<TextureAtlas>>,
    mut windows: Query<&mut Window>,
) {
    // 创建2d镜头
    commands.spawn(Camera2dBundle::default());

    // 获取当前窗口
    let window = windows.single_mut();
    let win_w = window.height();
    let win_h = window.height();

    //  添加 WinSize 资源
    commands.insert_resource(resource::WinSize { w: win_w, h: win_h });

    // 添加 GameTextures
    let images = resource::GameImages {
        broad: asset_server.load(resource::asset::IMAGE_MEI),
    };

    let sounds = resource::GameSound {
        click: asset_server.load(resource::asset::SOUND_CLICK),
    };

    let fonts = resource::GameFont {
        wenkai: asset_server.load(resource::asset::FONT_WENKAI)
    };

    // 棋盘
    commands.spawn(SpriteBundle {
        texture: images.broad.clone(),
        sprite: Sprite {
            custom_size: Some(Vec2 { x: win_w, y: win_h }),
            ..Default::default()
        },
        transform: Transform {
            translation: Vec3 {
                x: 120.,
                y: 120.,
                z: 11.,
            },
            ..Default::default()
        },
        ..Default::default()
    });


    // 字体引入
    let text_style = TextStyle {
        font: fonts.wenkai.clone(),
        font_size: 32.,
        color: Color::MIDNIGHT_BLUE,
    };
    let text_alignment = TextAlignment::Center;

    // 分数展示控件
    commands.spawn((
        Text2dBundle {
            text: Text::from_section("SCORE:0", text_style).with_alignment(text_alignment),
            transform: Transform {
                translation: Vec3 {
                    x: 0.,
                    y: win_h / 2. - 20.,
                    z: 11.,
                },
                ..Default::default()
            },
            ..Default::default()
        },
        component::DisplayScore,
    ));


    commands.insert_resource(sounds);
    commands.insert_resource(images);
}
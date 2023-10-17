use bevy::{math::Vec3Swizzles, prelude::*, sprite::collide_aabb::collide, utils::HashSet};


mod component;
mod resource;


fn main() {
    // add_startup_system 启动生命周期时只运行一次 ，
    // add_system 每帧都会被调用方法
    App::new()
        .add_state::<resource::GameState>()
        .insert_resource(ClearColor(Color::rgb(210., 210., 210.)))
        .add_plugins(DefaultPlugins.set(WindowPlugin {
            primary_window: Some(Window {
                title: "中国象棋".into(),
                resolution: (1000., 700.).into(),
                // position: WindowPosition::At(IVec2::new(2282, 0)),
                ..Window::default()
            }),
            ..WindowPlugin::default()
        }))
        .add_systems(Startup, setup_system)
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
        broad: asset_server.load(resource::asset::IMAGE_BROAD),
    };

    let sounds = resource::GameSound {
        click: asset_server.load(resource::asset::SOUND_CLICK),
    };


    // 棋盘背景图片
    commands.spawn(SpriteBundle {
        texture: images.broad.clone(),
        sprite: Sprite {
            custom_size: Some(Vec2 { x: 521., y: 577. }),
            ..Default::default()
        },
        transform: Transform::from_scale(Vec3::new(1.0, 1.0, 1.0)),
        ..Default::default()
    });

    commands.insert_resource(sounds);
    commands.insert_resource(images);
}
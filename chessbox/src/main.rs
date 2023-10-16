use bevy::{prelude::{App, DefaultPlugins, PluginGroup, default, Update, Msaa, Commands, Camera2dBundle, ResMut}, window::{WindowPlugin, PresentMode, Window}};

mod components;
mod asset;
use asset::{AssetPlugin, FONT_HANDLE, PIECE_HANDLE};
use components::GameTimers;


fn main() {
    App::new()
    .add_plugins(DefaultPlugins.set(WindowPlugin {
        primary_window: Some(Window {
            title: "中国象棋".into(),
            resolution: (800., 800.).into(),
            present_mode: PresentMode::AutoVsync,
            fit_canvas_to_parent: true,
            prevent_default_event_handling: false,
            ..default()
        }),
        ..default()
    }))
    .add_plugins(AssetPlugin)
    .insert_resource(Msaa::Sample8)
    .init_resource::<GameTimers>()
    .add_systems(Update, hello_world_system)
    .run();

}

fn setup_camera(mut commands: Commands, mut countdown: ResMut<GameTimers>) {
	commands.spawn(Camera2dBundle::default());
	countdown.black.pause();
}

fn hello_world_system() {
    println!("hello world");
}

use bevy::{prelude::{App, Camera2dBundle, Commands, default, DefaultPlugins, Msaa, PluginGroup, ResMut, Update}, window::{PresentMode, Window, WindowPlugin}};
use bevy::diagnostic::{FrameTimeDiagnosticsPlugin, LogDiagnosticsPlugin};
use bevy::app::AppExit;
use bevy::prelude::EventWriter;

fn exit_system(mut exit: EventWriter<AppExit>) {
    exit.send(AppExit);
}

use asset::AssetPlugin;
use components::GameTimers;
use asset::PIECE_HANDLE;

mod components;
mod asset;

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
        .add_plugins(LogDiagnosticsPlugin::default())
        .add_plugins(FrameTimeDiagnosticsPlugin::default())
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
    println!("{} {}", PIECE_HANDLE.is_weak(), PIECE_HANDLE.is_strong());
    println!("hello world");
}

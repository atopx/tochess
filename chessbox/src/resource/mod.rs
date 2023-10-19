pub(crate) mod asset;
pub mod timer;

use bevy::prelude::{AudioSource, Component, Font, Handle, Image, Resource, States, TextureAtlas, Timer, TimerMode, Entity};

pub struct Player(Entity);

#[derive(Resource)]
pub struct Game {}


/// 游戏窗口大小资源
#[derive(Resource)]
pub struct WinSize {
    pub w: f32,
    pub h: f32,
}


/// 游戏图像资源
#[derive(Resource)]
pub struct GameImages {
    pub broad: Handle<Image>,
}

#[derive(Resource)]
pub struct GameSound {
    pub click: Handle<AudioSource>,
}

#[derive(Resource)]
pub struct GameFont {
    pub wenkai: Handle<Font>,
}


/// 游戏状态
#[derive(Debug, Clone, Copy, Default, Eq, PartialEq, Hash, States)]
pub enum GameState {
    /// 等待
    #[default]
    PENDING,
    /// 游戏中
    RUNNING,
    /// 暂停
    PAUSED,
}

/// 玩家状态
#[derive(Resource)]
pub struct PlayerState {
    /// 就绪
    pub pending: bool,
}

impl Default for PlayerState {
    fn default() -> Self {
        Self {
            pending: false,
        }
    }
}


impl PlayerState {
    /// 被命中
    pub fn swap(&mut self) {
        self.pending = !self.pending;
    }
}

/// 游戏数据
#[derive(Resource)]
pub struct GameData {
    score: u32,
}

